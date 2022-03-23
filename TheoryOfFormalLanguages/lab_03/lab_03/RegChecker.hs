{-# LANGUAGE MultiWayIf #-}
module Parsing where

import Data.Char ( isDigit, isSpace, isAlphaNum, isAlpha, isLower, isUpper, digitToInt )
import Data.Maybe as Maybe
import qualified Data.Map as Map

import Text.Printf

import Control.Applicative  --( Alternative(empty, (<|>), many, ) )
import Data.List as List
import System.Directory

import qualified Data.Set as Set
import qualified Data.Maybe as Maybe


filterSpaces :: String -> String
filterSpaces = filter (not . isSpace)

newtype Parser a = P (String -> Maybe(a, String))
parse :: Parser a -> String -> Maybe(a, String)
parse (P p) = p

item :: Parser Char
item = P (\x -> case x of
                [] -> Nothing
                (x : xs) -> Just(x, xs))

instance Functor Parser where
   -- fmap :: (a -> b) -> Parser a -> Parser b
   fmap g p = P (\x -> case parse p x of
                            Nothing   -> Nothing
                            Just(v, xs) -> Just(g v, xs))

instance Applicative Parser where
   -- pure :: a -> Parser a
   pure v = P (\x -> Just(v, x))

   -- <*> :: Parser (a -> b) -> Parser a -> Parser b
   pg <*> px = P (\x -> case parse pg x of
                             Nothing   -> Nothing
                             Just(g, xs) -> parse(fmap g px) xs)

instance Monad Parser where
   -- (>>=) :: Parser a -> (a -> Parser b) -> Parser b
   p >>= f = P (\x -> case parse p x of
                           Nothing       -> Nothing
                           Just(val, xs) -> parse (f val) xs)

instance Alternative Parser where
   -- empty :: Parser a
   empty = P (const Nothing)

   -- (<|>) :: Parser a -> Parser a -> Parser a
   p <|> q = P (\x -> case parse p x of
                           Nothing       -> parse q x
                           Just (val, xs) -> Just (val, xs))

satisfy :: (Char -> Bool) -> Parser Char
satisfy p = item >>=
            (\x -> if p x
                   then pure x
                   else empty) --- wired


letter :: Parser String
letter = some $ satisfy isLower

lowerLetter :: Parser Char
lowerLetter = satisfy isLower

lowerLetterStr :: Parser String
lowerLetterStr = fmap List.singleton lowerLetter

letterTerm :: Parser Term
letterTerm = Terminal <$> lowerLetter

upperLetter :: Parser Char
upperLetter = satisfy isUpper

digit :: Parser Char
digit = satisfy isDigit

char :: Char -> Parser Char
char c = satisfy (== c)


nterm :: Parser String
nterm = do
    x1 <- upperLetter
    x2 <- many digit
    return $  x1 : x2

term :: Parser Term
term = Nterm <$> nterm <|> letterTerm


data Rule = Rule{
    left :: String,
    right :: [Term]
} deriving(Show)


data Term = Nterm String | Terminal Char deriving (Show)

instance Eq Term where
  (Nterm a) == (Nterm b) = a == b
  (Terminal a) == (Terminal b) = a == b
  _ == _ = False


instance Ord Term where
    compare (Nterm a) (Nterm b) = compare a b
    compare (Terminal a) (Terminal b) = compare [a] [b]
    compare (Terminal a) (Nterm b) =  compare [a] b
    compare (Nterm a) (Terminal b) = compare a [b]


data CFG = CFG{
    cfgRules :: Map.Map String [Rule]
} deriving (Show)

makeCFG :: [Rule] -> CFG
makeCFG rules = CFG{cfgRules= Map.fromList  (wrapKV <$> List.groupBy (\x y -> left x == left y) sortedRules)}
    where wrapKV rules = (left $ head rules , rules)
          sortedRules = List.sortBy (\x y -> compare (left x) (left y)) rules


rule = do
        x1 <- nterm
        char '-'
        char '>'
        x2 <- letterTerm
        x3 <- some term
        return Rule{left = x1, right = x2 : x3}
       <|>
       do       
        x1 <- nterm
        char '-'
        char '>'
        x2 <- letterTerm
        return Rule{left = x1, right = [x2]}


ntermsFromCfg :: CFG -> Set.Set String
ntermsFromCfg cfg = Map.keysSet $ cfgRules cfg

ntermsFromCfg2 :: CFG -> Set.Set Term
ntermsFromCfg2 cfg = Set.map Nterm (Map.keysSet $ cfgRules cfg)

isRecursive :: Rule -> Bool
isRecursive r = isJust $ List.find p (right r)
    where p t = case t of
            Nterm s -> s == left r
            Terminal c -> False

isNterm :: Term -> Bool
isNterm (Nterm _) = True
isNterm (Terminal _) = False

isTerminal = not . isNterm


fromNterm :: Term -> String
fromNterm (Nterm s) = s
fromNterm (Terminal _) = error "PANIC"


isRightLinear :: Rule -> Bool
isRightLinear r = length rightPart <= 1
    where rightPart = dropWhile (not . isNterm) (right r)


ntermsCount :: Rule -> Int
ntermsCount r = sum $ fromEnum. isNterm <$> right r


isConst :: Rule -> Bool
isConst r = not $ or $ isNterm <$> right r


ntermsFrom cfg = filter isNterm . concat $ right <$> rules
    where rules = concat $ snd <$> (Map.toList (cfgRules cfg))


remover :: Set.Set String -> CFG -> Set.Set String
remover subset cfg
    | (Set.size subset) == (Set.size $ removeHelper subset) = subset
    | otherwise = remover (removeHelper subset) cfg

    where removeHelper set = Set.filter p set
          rules = cfgRules cfg
          p x = and $ (`Set.member` subset) <$> (ntermsFromRules $ rules Map.! x)


ntermsFromRules :: [Rule] -> [String]
ntermsFromRules rules = fromNterm <$> (filter isNterm . concat)  (right <$> rules)

regularSubset cfg = remover subset cfg
    where nterms = fromNterm <$> ntermsFrom cfg
          isRegular ntName = all isRightLinear (cfgRules cfg Map.! ntName)
          subset = Set.fromList $ filter isRegular nterms


data PumpTree = Node [PumpTree] Term | Leaf Term | None deriving (Show, Eq)

addChildToNode :: PumpTree -> PumpTree -> PumpTree
addChildToNode a None = a
addChildToNode (Node xs t) b = Node (xs ++ [b]) t
addChildToNode _ _ = undefined


makePumpTreeRun :: CFG -> Term -> (PumpTree, Bool)
makePumpTreeRun cfg ter  = makePumpTree cfg ter Set.empty ter


makeNode (Terminal t) = Leaf (Terminal t)
makeNode (Nterm t) = Node [] (Nterm t)


rotatePumpTree (Node xs t) = Node (reverse (rotatePumpTree <$> xs)) t
rotatePumpTree (Leaf t) = Leaf t
rotatePumpTree None = None


retrievePumping :: PumpTree -> [Term]
retrievePumping (Node xs t) | length xs > 0 = concat $ retrievePumping <$> xs
                            | otherwise = [t]
retrievePumping (Leaf t) = [t]
retrievePumping None = []


makePumpTree :: CFG -> Term -> Set.Set Term -> Term -> (PumpTree, Bool)
makePumpTree cfg startnTerm visited ter | isTerminal ter = (Leaf ter , False)
                                        | Set.member ter visited = if ter == startnTerm then (Node [] ter, True) else (None, False)
                                        | otherwise = lookupRules $ rules ter

    where rules (Nterm t) = cfgRules cfg Map.! t
          rules (Terminal t) = []

          lookupRules (r : rs) | snd $ lookupRighsSideRun $ right r = (Node (fst $ lookupRighsSideRun $ right r) ter, True)
                               | otherwise = lookupRules rs

          lookupRules [] = (None, False)

          stepTo x = makePumpTree cfg startnTerm (Set.insert ter visited) x

          lookupRighsSideRun xs = lookupRighsSide xs False []

          lookupRighsSide (x : xs) True children = lookupRighsSide xs True ((makeNode x) : children)
          lookupRighsSide [] True children = (children, True)


          lookupRighsSide (x : xs) False children | fst (stepTo x) == None = (children, False)
                                                  | snd $ stepTo x  = lookupRighsSide xs True (fst (stepTo x) : children)
                                                  | otherwise = lookupRighsSide xs False (fst (stepTo x) : children)

          lookupRighsSide [] False children = (children, False)


termToStr (Nterm t) = t
termToStr (Terminal t) = [t]

termListToStr xs = concat $ termToStr <$> xs

phi1InPhi2P f1 f2 cfg = helper f1 f2
    where helper (x1 : xs1) (x2 : xs2) | x1 == x2 = helper xs1 xs2
          helper [] (x2:xs2) = True
          helper f1_suff [] = length f2 /= 0 && helper f1_suff f2
          helper _ (Terminal t : xs2) = False
          helper f1_suff f2_suff = foldl (\x y -> x || (helper f1_suff (right y ++ (drop 1 f2_suff)))) False (fromMaybe [] (cfgRules cfg Map.!? (termToStr (head f2_suff))))


testPump = retrievePumping $ rotatePumpTree $ fst $ makePumpTreeRun testCFG (Nterm "S")


splitPump :: [Term] -> Term -> ([Term], [Term])
splitPump xs e = splitPumpGo xs []
    where splitPumpGo (l : ls) rs | l /= e = splitPumpGo ls (l : rs)
                                  | otherwise = (reverse rs, ls)
          splitPumpGo [] rs = (reverse rs, [])


rulesForTerm :: CFG -> String -> [Rule]
rulesForTerm cfg term = fromMaybe [] $ cfgRules cfg Map.!? term
rulesForTermT cfg term = rulesForTerm cfg  (termToStr term)


unfold :: Term -> CFG -> [[Term]]
unfold ter cfg = runn 1000 [[ter]]
    where runn ticks words
            -- | ticks == 0 = words
            | or $ isTerminalString <$> words = filter isTerminalString words
            | otherwise = runn (ticks - 1) $ concat (unfoldWord [] <$> words)

          unfoldWord ls (Nterm t : ts)  = ((\x -> ls ++ x ++ ts) . right <$> rules t) ++ unfoldWord (ls ++ [Nterm t]) ts
          unfoldWord ls (Terminal t : ts)  = unfoldWord (ls ++ [Terminal t]) ts

          unfoldWord ls [] | isTerminalString ls = [ls]
                           | otherwise = []

          isTerminalString :: [Term] -> Bool
          isTerminalString word = and (isTerminal <$> word)

          rules t = cfgRules cfg Map.! t


makeTerm :: [Char] -> Term
makeTerm t | length t == 1 && isLower (t !! 1) = Terminal (t !! 1)
           | length t >= 1 && isUpper (t !! 1) = Nterm t
           | otherwise = error "empty term"


makeTerm1 :: Char -> Term
makeTerm1 t | isLower t = Terminal t
            | isUpper t = Nterm [t]
            | otherwise = error "empty term"

rulesFromStr :: [Char] -> [Term]
rulesFromStr xs = makeTerm1 <$> xs


checkPumps :: CFG -> Set.Set String -> ([Term], [Term], [Term])
checkPumps cfg regSubset = checkPumpsHelper ([], [], List.map Nterm (Set.toList regSubset)) nonRegNterms
    where checkPumpsHelper triple (x : xs) = checkPumpsHelper (appendTrip (checkPump cfg regSubset x) triple) xs
          checkPumpsHelper triple [] = triple

          appendTrip (x1, x2, x3) (y1, y2, y3) = (x1 ++ y1, x2 ++ y2, x3 ++ y3)
          nonRegNterms = Nterm <$> (Set.toList $ ntermsFromCfg cfg Set.\\ regSubset)


ntermInRightParts :: CFG -> Term -> [Term]
ntermInRightParts cfg t = Nterm <$> ntermsFromRules ((rulesForTerm cfg . fromNterm) t)


closureRunner :: CFG -> [Term] -> [Term] -> Set.Set Term  -> ([Term], [Term])
closureRunner cfg regular mbRegular unknownT
        | isUpdated = closureRunner cfg newReg newMbReg newUnknown
        | otherwise = (regular, mbRegular)

    where newPair = closure cfg regular mbRegular (Set.toList unknownT)
          newReg = fst newPair
          newMbReg = snd newPair
          isUpdated = length newReg /= length regular || length newMbReg /= length mbRegular
          newUnknown = (ntermsFromCfg2 cfg) Set.\\ (Set.fromList regular) Set.\\ (Set.fromList mbRegular)


closure :: CFG -> [Term] -> [Term] -> [Term] -> ([Term], [Term])
closure cfg regular mbRegular (t : ts)
           | checkRegularity t = closure cfg (t : regular) mbRegular ts
           | checkMbRegularity t = closure cfg  regular (t : mbRegular) ts
           | otherwise = closure cfg regular mbRegular ts

    where checkMbRegularity t = recCheck checkNTMbRegularity $ t
          checkRegularity t = recCheck checkNTRegularity $ t
          checkNTMbRegularity x =  (x `List.elem` mbRegular) || (x `List.elem` regular)
          checkNTRegularity = (`List.elem` regular)
          recCheck f t = all f $ ntermInRightParts cfg t

closure cfg regular mbRegular [] = (regular, mbRegular)


isRegular :: Set.Set String -> [Term] -> Bool
isRegular regSubset xs = helper xs
    where helper (Nterm t : xs) = Set.member t regSubset && helper xs
          helper (Terminal t : xs) = helper xs
          helper [] = True


checkPump :: CFG -> Set.Set String ->  Term -> ([Term], [Term], [Term])
checkPump cfg regSubset nterm
    | not $ snd pump = ([], [], [])
    | not isf2Reg || not checkIsInF2P = ([nterm], [], [])
    | checkUnfolds && Set.member (fromNterm nterm) regSubset = ([], [], [nterm])
    | checkUnfolds = ([], [nterm], [])
    | otherwise = ([], [], [])-- nterm ++ show f1 ++ show f2 ++ show (unfold nterm cfg)

    where pump = makePumpTreeRun cfg nterm
          split = (splitPump $ retrievePumping $ fst pump) nterm
          f1 = reverse $ fst split
          f2 = reverse $ snd split
          isf2Reg = isRegular regSubset f2
          checkIsInF2P = phi1InPhi2P f1 f2 cfg
          checkUnfolds = all (\x -> phi1InPhi2P x f2 cfg) (unfold nterm cfg)


fst3 (x, _, _) = x
snd3 (_, x, _) = x
trd3 (_, _, x) = x


testsDir = "./tests"

tests :: IO [FilePath]
tests = List.map ((testsDir ++ "/") ++) <$> listDirectory testsDir

main :: IO ()
main = do
    t <- tests
    mapM runTest t
    return()


runTest :: FilePath -> IO()
runTest file = do
   putStrLn  $ "Run test" ++ file ++ "\n"
   
   cfgM <- readRules file
   if isNothing cfgM
      then do print "Syntax error"
              return ()
      else analyze $ Maybe.fromJust cfgM


readRules :: FilePath -> IO (Maybe CFG)
readRules file = do
    fileData <- readFile file
    let reqLines = filterSpaces <$> lines fileData
        rulesM = parse rule <$> reqLines
        cfg =  makeCFG $ fst . fromJust <$> rulesM

    if any Maybe.isNothing rulesM || any (/= 0)  (length . snd . fromJust <$> rulesM)
        then return Maybe.Nothing
        else return $ Just cfg


analyze :: CFG -> IO ()
analyze cfg = do
    putStrLn "Parsed ..."
    let regSubset = regularSubset cfg
        triple = checkPumps cfg regSubset
        susp = fst3 triple
        mbReg = snd3 triple
        reg = trd3 triple
        closure = closureRunner cfg reg mbReg (ntermsFromCfg2 cfg)
        newReg = (Set.fromList $ fst closure)
        newMbReg = (Set.fromList $ snd closure)
        unknown = (ntermsFromCfg2 cfg) Set.\\ newReg Set.\\ newMbReg
        susp' = (Set.fromList susp) Set.\\ newReg Set.\\ newMbReg
    
    -- print triple

    -- print "Regular subset"
    -- print $ regSubset

    putStrLn $ (regStatus susp' newReg newMbReg unknown)

    putStrLn "Regular NotTerminals:"
    putStrLn $ "\t" ++ show newReg

    putStrLn "Possibly regular NotTerminals:"
    putStrLn $ "\t" ++ show newMbReg

    putStrLn "Suspicious NotTerminals:"
    putStrLn $ "\t" ++ show susp

    putStrLn "----------"


regStatus susp regularNT mbRegular unknown
    | length susp > 0 = "Language is suspicious"
    | length unknown > 0 = "Language regularity could not be determined"
    | length mbRegular > 0 = "Language is possibly regular"
    | length regularNT > 0 = "Language is regular"


testCFG :: CFG
testCFG = makeCFG rulesTest1

-- rulesTest1 :: [Rule]
-- rulesTest1 = [
--              Rule {left = "S", right = rulesFromStr "aSA"},
--              Rule {left = "S", right = rulesFromStr "aA"},
--              Rule {left = "A", right = rulesFromStr "a"}
--             ]


rulesTest1 :: [Rule]
rulesTest1 = [
                Rule {left = "S", right = rulesFromStr "aA"},
                Rule {left = "S", right = rulesFromStr "aBb"},
                Rule {left = "S", right = rulesFromStr "bBa"},
                Rule {left = "B", right = rulesFromStr "aaSS"},
                Rule {left = "S", right = rulesFromStr "a"},
                Rule {left = "A", right = rulesFromStr "babaAba"},
                Rule {left = "A", right = rulesFromStr "ba"}
            ]


testRs = regularSubset testCFG

module RegexEqSystemSolver where
import Parsing
import Data.Map as Map
import Data.Maybe
import Data.Set as Set
import Control.Applicative
import Data.List as List


expressVar' :: RegexEq' -> RegexEq'
expressVar' r | coeff' alpha == "" = R{variable'= v, linComb = betha}
              | otherwise = R{variable'= v, linComb = Map.map (klenneStar alpha <|*>) betha}
   where v = variable' r
         alpha =  Map.findWithDefault (C "" "") v (linComb r)
         betha = filterWithKey (\x _ -> x /= v) $ linComb r


substitude :: RegexEq' -> RegexEq' -> RegexEq'
substitude a b
   | Map.member varA linCombB = R{variable' = variable' b, linComb = m}
   | otherwise = b
   where varA = variable' a
         linCombB = linComb b
         m = linCombB <|++> Map.map (alpha' <|*>) (linComb . expressVar' $ a)
         x <|++> y = Map.delete varA (unionWith combiner x y)
         alpha' = Map.findWithDefault (C "" "") varA linCombB
         combiner l r = r <+> l

applyToCoef :: (String -> String) -> CoeffVar' -> CoeffVar'
applyToCoef f c = C (variable c) (f $ coeff' c)


wrapInParen :: CoeffVar' -> CoeffVar'
wrapInParen = applyToCoef wrap
    where wrap s = "(" ++ s ++ ")"


isInParrens :: CoeffVar' -> Bool
isInParrens t = f == '(' && (l == ')' || l == '*')
   where f = head . coeff' $ t
         l = last . coeff' $ t


klenneStar :: CoeffVar' -> CoeffVar'
klenneStar t  
   | (head . coeff' $ t) == '(' && (last . coeff' $ t) == ')' = applyToCoef (++"*") t
   | (length . coeff' $ t) == 1 = applyToCoef (++"*") t
   | otherwise = applyToCoef (++"*") $ wrapInParen t


(<+>) :: CoeffVar' -> CoeffVar' -> CoeffVar'
a <+> b = if variable a == variable b
          then C (variable a) (coeff' a ++ "+" ++ coeff' b)
          else undefined


(<|*>) :: CoeffVar' -> CoeffVar' -> CoeffVar'
a <|*> b = C {variable = variable b, coeff' = coeff' aa ++ coeff' bb}
   where aa | (length . coeff' $ a) == 1 = a 
            | isInParrens a = a 
            | otherwise = wrapInParen a

         bb | (length . coeff' $ b) == 1 = b 
            | isInParrens b = b
            | otherwise = wrapInParen b


retrieve :: RegexEq' -> [Char]
retrieve R{variable' = v, linComb = lc} = v ++ "=" ++ coeff' (fromMaybe (C "" "") (lc Map.!? ""))


-- solve :: [RegexEq'] -> [[Char]]
-- solve xs = retrieve . expressVar' <$> (backward . forward $ xs)

solve :: [RegexEq'] -> [[Char]]
solve xs = retrieve . expressVar' <$> (forward . forward $ xs)

a <|-> b = b `substitude` a
infixl 5 <|->

forward :: [RegexEq'] -> [RegexEq']
forward xs = forwardGo [head xs] (List.drop 1 xs)
    where forwardGo h [] = h
          forwardGo h (x : xs) = forwardGo ( substFrom h x : h) xs
          substFrom h x = List.foldl1 (<|->) (x : h)


-- backward :: [RegexEq'] -> [RegexEq']
-- backward xs = reverse $ scanl1 substitude (reverse xs)


checkVarsInSystem ::  [RegexEq'] -> Bool 
checkVarsInSystem eqSys = not $ Set.isSubsetOf varsRight varsLeft
   where varsLeft = Set.fromList $  variable' <$> eqSys 
         varsRight = List.foldr Set.union Set.empty $ Map.keysSet . Map.delete "" . linComb <$> eqSys


solveSystemFrom :: FilePath -> IO ()
solveSystemFrom file = do
    fileData <- readFile file
    putStrLn $ "Run test " ++ file
    let reqLines = filterSpaces <$> lines fileData
        equations = parse equation <$> reqLines
        equations' = fst . fromJust <$> equations
        system = fst . fromJust <$> equations
        runner
            | isJust(find (== Nothing ) equations) || (not $ and ((== 0) . length . snd . fromJust <$> equations)) = do
                putStrLn $ id "SyntaxError"
                return()
            | checkVarsInSystem system = do
                putStrLn $ id "There are vars with no Equation for them"
                return()
            | otherwise = do
                putStrLn $ id "Solution"
                mapM print (solve system)
                return()
    runner

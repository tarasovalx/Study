module RGtoRegex where

import RegexEqSystemSolver
import Parsing
import Data.Maybe
import Data.Map as Map
import Data.List as List
import Data.Set as Set

buildSystem :: [RegexEq'] -> [RegexEq']
buildSystem = Prelude.map snd . Map.toList . Map.fromListWith (<++>) . Prelude.map (\x -> (variable' x, x))
    where a <++> b = R{variable' = variable' a, linComb = Map.unionWith (<+>) (linComb a) (linComb b)}


-- checkForTraps :: [RegexEq'] -> Bool
-- checkForTraps rules = or $ checkForRule <$> rules
--     where checkForRule r = or $ checkForVar  (variable' r) . fst <$>  Map.toList (linComb r)
--           checkForVar eqV lcV | eqV == lcV = False
--                               | lcV == "" = False
--                               | otherwise = or $ findInLinComb eqV <$> rulesForVar lcV

--           findInLinComb vName r = Map.member vName (linComb r)
--           rulesForVar v = List.filter ((==v) . variable') rules


checkForTraps' :: [RegexEq'] -> Bool
checkForTraps' rules  | suspiciousRules == [] = False
                      | otherwise = checkTrapsForVar' rules  (Set.fromList suspiciousRules) 
    where suspiciousRules = variable' <$> List.filter (Map.notMember "" . linComb)  rules


checkTrapsForVar' :: [RegexEq'] -> Set.Set String -> Bool
checkTrapsForVar' system suspiciousRules = and $ (helper Set.empty) <$> ((systemMap Map.!)  <$> (Set.toList  suspiciousRules))
    where helper visited rule
                              | Set.notMember (variable' rule) suspiciousRules = False
                              | Set.member (variable' rule) visited = True
                              | otherwise =  and $ helper (Set.insert (variable' rule) visited) <$> (f rule) ---

          systemMap = Map.fromList ((\x -> (variable' x , x)) <$> system) 
          g eq = fst <$> (Map.toList . Map.filterWithKey (\k _ -> k /= "")) (linComb eq)
          f eq = (systemMap Map.!)  <$> (g eq)

rgToRegexFrom :: FilePath -> IO ()
rgToRegexFrom file = do
    fileData <- readFile file
    putStrLn $ "Run test " ++ file
    let reqLines = filterSpaces <$> lines fileData
        rules = parse rule <$> reqLines
        rules' = fst . fromJust <$> rules
        system = buildSystem $ fst . fromJust <$> rules
        runner
            | isJust(find (== Nothing ) rules) || (not $ and ((== 0) . length . snd . fromJust <$> rules)) = do
                putStrLn $ id "SyntaxError"
                return()
            | checkVarsInSystem system = do
                putStrLn $ id "There are vars with no Equation for them"
                return()
            | checkForTraps' system = do
                -- print system
                putStrLn $ id "There are traps in grammar"
                return()
            | otherwise = do
                -- print system
                putStrLn $ id "Solution"
                mapM print (solve system)
                return()
    runner

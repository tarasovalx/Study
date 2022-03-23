module Runner where
import RGtoRegex
import RegexEqSystemSolver
import System.Directory
import Data.List as List


tests2Dir = "./tests2"
tests3Dir = "./tests3"

tests2 :: IO [FilePath]
tests2 = List.map ((tests2Dir ++ "/") ++) <$> listDirectory tests2Dir


tests3 :: IO [FilePath]
tests3 = List.map ((tests3Dir ++ "/") ++) <$> listDirectory tests3Dir

main :: IO ()
main = do
    putStrLn "SystemSolver test"
    t2 <- tests2
    mapM solveSystemFrom t2

    putStrLn ""

    putStrLn "RG to Regex test"
    t3 <- tests3
    mapM rgToRegexFrom t3
    --mapM rgToRegexFrom ["./tests3/test7"]
    return()

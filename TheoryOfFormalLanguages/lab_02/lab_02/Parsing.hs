module Parsing where

import Data.Char ( isDigit, isSpace, isAlphaNum, isAlpha, isLower, isUpper, digitToInt )
import Data.Maybe ( Maybe )
import qualified Data.Map as Map

import Text.Printf

import Control.Applicative 
import Data.List

newtype Parser a = P (String -> Maybe(a, String))
parse :: Parser a -> String -> Maybe(a, String)
parse (P p) = p

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
                           Just(val, xs) -> Just(val, xs))

item :: Parser Char
item = P (\x -> case x of
                [] -> Nothing
                (x : xs) -> Just(x, xs))

satisfy :: (Char -> Bool) -> Parser Char
satisfy p = item >>=
            (\x -> if p x
                   then pure x
                   else empty)


letter :: Parser String
letter = some $ satisfy isLower

lowerLetter :: Parser Char
lowerLetter = satisfy isLower

char :: Char -> Parser Char
char c = satisfy (== c)

var :: Parser String
var = some (satisfy isUpper)

expr :: Parser String
expr = some lowerLetter

altRegex :: Parser String
altRegex = do
           x1 <- expr
           char '|'
           printf "(%s|%s)" x1 <$> altRegex
           <|>
           expr


regex :: Parser String
regex = do
        char '('
        x1 <- expr
        char '|'
        x2 <- altRegex
        char ')'
        return (printf "(%s|%s)" x1 x2)
        <|>
        expr

type Varname = String

data CoeffVar' = C{
   variable :: Varname,
   coeff' :: String
} deriving(Show, Eq)

C v c = C{variable = v, coeff'=c}

data RegexEq' = R{
   variable' :: Varname,
   linComb :: Map.Map Varname CoeffVar'
} deriving(Show, Eq)


regexVar :: Parser CoeffVar'
regexVar = do
            x1 <- regex
            x2 <- var
            return $ C x2 x1


space :: Parser ()
space = do many (satisfy isSpace)
           return ()


zeroOrOne :: Parser a -> Parser [a]
zeroOrOne p = fmap (: []) p <|> P(\x -> Just([],  x))


equation :: Parser RegexEq'
equation = do
           x1 <- var
           char '='
           x2 <- zeroOrOne (do xx1 <- regexVar
                               xx2 <- many (do char '+'
                                               regexVar)
                               char '+'
                               return (xx1 : xx2))
           x3 <- regex
           return R {variable' = x1, linComb = Map.fromList (map (\x -> (variable x, x)) (C "" x3 : concat x2))}
           <|>
           do
           x1 <- var
           char '='
           x2 <- regex
           return R {variable' = x1, linComb = Map.fromList [("", C "" x2)]}


filterSpaces :: String -> String
filterSpaces = filter (not . isSpace)

--- Parse RG
nterm :: Parser String
nterm = some (satisfy isUpper)

rule :: Parser RegexEq'
rule = do x1 <- nterm
          char '-'
          char '>'
          x2 <- letter
          x3 <- nterm
          return R{variable' = x1, linComb = Map.fromList [(x3 , C x3 x2)]}
       <|>
       do x1 <- nterm
          char '-'
          char '>'
          x2 <- letter
          return R{variable' = x1, linComb = Map.fromList [("" , C "" x2)]}
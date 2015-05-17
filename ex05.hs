module Main where

import Data.List

digits :: [Char]
digits = ['1'..'9']

ops :: [Char]
ops = ['+', '-']

solve :: [String]
solve = [ d | d <- inits digits ]

main :: IO ()
main = mapM_ print solve

module Main where

import Control.Applicative
import Data.List
import Data.Function
import System.Environment

sort' :: [Int] -> [Int]
sort' = sortBy (compare `on` (show :: Int -> String))

solve :: [Int] -> Int
solve = read . concat . map show . reverse . sort'

main :: IO ()
main = do
  nums <- map read <$> getArgs
  print $ solve nums

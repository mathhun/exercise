module Main where

main :: IO ()
main = print $ zip (take 10 [1..]) (take 10 ['a'..])

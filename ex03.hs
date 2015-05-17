module Main where

fibonacci :: [Integer]
fibonacci = 0:1:zipWith (+) fibonacci (tail fibonacci)

fib :: Int -> [Integer]
fib n = take n fibonacci

main :: IO ()
main = print $ fib 100

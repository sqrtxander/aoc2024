import System.IO (readFile)
import Data.List (sort, span)
import Data.Char (isSpace)

splitRow :: String -> (String, String)
splitRow row = let (left, rest) = break isSpace row
                   right = dropWhile isSpace rest
               in (left, right)

getNums :: [String] -> ([Int], [Int])
getNums rows = let (leftStr, rightStr) = unzip $ map splitRow rows
                   left = map read leftStr
                   right = map read rightStr
               in (left, right)

solve :: [String] -> Int
solve rows = let (left, right) = getNums rows
             in sum $ zipWith (\a b -> abs(a-b)) (sort left) (sort right)

main :: IO()
main = do
    rows <- lines <$> readFile "input.in"
    print $ solve rows

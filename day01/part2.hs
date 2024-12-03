import Data.Char (isSpace)
import qualified Data.IntMap as IntMap
import Data.List (sort, span)
import System.IO (readFile)

splitRow :: String -> (String, String)
splitRow row =
    let
        (left, rest) = break isSpace row
        right = dropWhile isSpace rest
    in
        (left, right)

getNums :: [String] -> ([Int], [Int])
getNums rows =
    let
        (leftStr, rightStr) = unzip $ map splitRow rows
        left = map read leftStr
        right = map read rightStr
    in
        (left, right)

countOccurrences :: [Int] -> IntMap.IntMap Int
countOccurrences = foldr (\n acc -> IntMap.insertWith (+) n 1 acc) IntMap.empty

solve :: [String] -> Int
solve rows =
    let
        (left, right) = getNums rows
        rightMap = countOccurrences right
    in
        sum $ map (\n -> n * IntMap.findWithDefault 0 n rightMap) left

main :: IO ()
main = do
    rows <- lines <$> readFile "input.in"
    print $ solve rows

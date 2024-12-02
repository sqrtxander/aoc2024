import System.IO (readFile)

getNumsRow :: String -> [Integer]
getNumsRow row = map read $ words row

getNums :: [String] -> [[Integer]]
getNums = map getNumsRow

isValidAsc :: [Integer] -> Bool
isValidAsc [] = True
isValidAsc [x] = True
isValidAsc (x : y : xs) = y - x >= 1 && y - x <= 3 && isValidAsc (y : xs)

isValidDsc :: [Integer] -> Bool
isValidDsc [] = True
isValidDsc [x] = True
isValidDsc (x : y : xs) = x - y >= 1 && x - y <= 3 && isValidDsc (y : xs)

isValid :: [Integer] -> Bool
isValid xs = isValidAsc xs || isValidDsc xs

solve :: [String] -> Int
solve rows = length $ filter isValid $ getNums rows

main :: IO ()
main = do
    rows <- lines <$> readFile "input.in"
    print $ solve rows

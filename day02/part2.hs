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

removeIdx :: Integer -> [a] -> [a]
removeIdx 0 (x : xs) = xs
removeIdx y (x : xs) = x : removeIdx (y - 1) xs

isValidArr :: [Integer] -> Bool
isValidArr xs = isValidAsc xs || isValidDsc xs

isValid :: [Integer] -> Bool
isValid xs = isValidArr xs || isValidHelper 0 xs

isValidHelper :: Integer -> [Integer] -> Bool
isValidHelper y xs
    | y >= fromIntegral (length xs) = False
    | otherwise = isValidArr (removeIdx y xs) || isValidHelper (y + 1) xs

solve :: [String] -> Int
solve rows = length $ filter isValid $ getNums rows

main :: IO ()
main = do
    rows <- lines <$> readFile "input.in"
    print $ solve rows

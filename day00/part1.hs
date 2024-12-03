import System.IO

solve :: [String] -> Integer
solve rows = 1

main :: IO ()
main = do
    rows <- lines <$> readFile "input.in"
    print $ solve rows

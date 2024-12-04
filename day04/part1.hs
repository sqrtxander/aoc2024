import qualified Data.Map as Map
import System.IO

type Grid = Map.Map (Integer, Integer) Char

parse :: [String] -> Grid
parse rows = Map.fromList [((x, y), z) | (y, row) <- zip [0 ..] rows, (x, z) <- zip [0 ..] row]

countWord :: String -> (Integer, Integer) -> Grid -> Integer
countWord s p g =
    let
        dirs = [(0, 1), (-1, 1), (-1, 0), (-1, -1), (0, -1), (1, 1), (1, 0), (1, -1)]
    in
        sum [countWordDirection s p dir g | dir <- dirs]

countWordDirection :: String -> (Integer, Integer) -> (Integer, Integer) -> Grid -> Integer
countWordDirection "" _ _ _ = 1
countWordDirection (z : zs) (x, y) (dx, dy) g =
    case Map.lookup (x, y) g of
        Just c | c == z -> countWordDirection zs (x + dx, y + dy) (dx, dy) g
        _ -> 0

solve :: [String] -> Integer
solve rows =
    let
        grid = parse rows
    in
        Map.foldrWithKey (\k _ acc -> acc + countWord "XMAS" k grid) 0 grid

main :: IO ()
main = do
    rows <- lines <$> readFile "input.in"
    print $ solve rows

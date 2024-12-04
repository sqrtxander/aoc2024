import qualified Data.Map as Map
import Data.Maybe (catMaybes, isJust)
import System.IO

type Grid = Map.Map (Integer, Integer) Char

parse :: [String] -> Grid
parse rows = Map.fromList [((x, y), z) | (y, row) <- zip [0 ..] rows, (x, z) <- zip [0 ..] row]

isXMas :: (Integer, Integer) -> Grid -> Bool
isXMas (x, y) g =
    case Map.lookup (x, y) g of
        Just c | c == 'A' -> isRightCount (x, y) g && isRightPermutation (x, y) g
        _ -> False

isRightCount :: (Integer, Integer) -> Grid -> Bool
isRightCount (x, y) g =
    let
        dirs = [(-1, -1), (-1, 1), (1, 1), (1, -1)]
        chars = catMaybes [Map.lookup (x + dx, y + dy) g | (dx, dy) <- dirs]
    in
        length (filter (== 'S') chars) == 2 && length (filter (== 'M') chars) == 2

isRightPermutation :: (Integer, Integer) -> Grid -> Bool
isRightPermutation (x, y) g =
    let
        tl = Map.lookup (x - 1, y - 1) g
        br = Map.lookup (x + 1, y + 1) g
    in
        isJust tl && isJust br && tl /= br

solve :: [String] -> Integer
solve rows =
    let
        grid = parse rows
    in
        Map.foldrWithKey (\k _ acc -> acc + if isXMas k grid then 1 else 0) 0 grid

main :: IO ()
main = do
    rows <- lines <$> readFile "input.in"
    print $ solve rows

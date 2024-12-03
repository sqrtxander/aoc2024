import Data.Char (isDigit)
import Data.List (isPrefixOf)
import System.IO

getMulParams :: String -> Bool -> [(Integer, Integer)]
getMulParams "" _ = []
getMulParams s True
    | "don't()" `isPrefixOf` s =
        let
            rest = drop (length "don't()") s
        in
            getMulParams rest False
    | "mul(" `isPrefixOf` s =
        let
            rest = drop (length "mul(") s
            digits = checkMulParams rest
        in
            case digits of
                Nothing -> getMulParams rest True
                Just nums -> nums : getMulParams rest True
    | otherwise = getMulParams (tail s) True
getMulParams s False
    | "do()" `isPrefixOf` s =
        let
            rest = drop (length "do()") s
        in
            getMulParams rest True
    | otherwise = getMulParams (tail s) False

checkMulParams :: String -> Maybe (Integer, Integer)
checkMulParams s =
    case parseNum s of
        Nothing -> Nothing
        Just (first, rest1) -> case rest1 of
            ',' : rest2 -> case parseNum rest2 of
                Nothing -> Nothing
                Just (second, rest3) -> case rest3 of
                    ')' : _ -> Just (first, second)
                    _ -> Nothing
            _ -> Nothing

parseNum :: String -> Maybe (Integer, String)
parseNum s =
    let
        (digits, rest) = span isDigit s
    in
        if null digits then Nothing else Just (read digits, rest)

solve :: String -> Integer
solve contents = sum $ map (uncurry (*)) $ getMulParams contents True

main :: IO ()
main = do
    contents <- readFile "input.in"
    print $ solve contents

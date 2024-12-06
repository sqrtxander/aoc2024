import qualified Data.IntMap as IntMap
import Data.List (isPrefixOf, sortBy)
import GHC.Exts (sortWith)
import Text.Parsec
import Text.Parsec.String (Parser)

ruleParser :: Parser (Integer, Integer)
ruleParser = do
    before <- read <$> many1 digit
    char '|'
    after <- read <$> many1 digit
    return (after, before)

updateParser :: Parser [Integer]
updateParser = do
    (read <$> many1 digit) `sepBy` char ','

inputParser :: Parser (IntMap.IntMap [Integer], [[Integer]])
inputParser = do
    rulePairs <- many $ try (ruleParser <* newline)
    newline
    updates <- updateParser `endBy` newline
    eof
    return (unionise rulePairs, updates)

unionise :: [(Integer, Integer)] -> IntMap.IntMap [Integer]
unionise = IntMap.fromListWith (++) . map (\(a, b) -> (fromIntegral a, [fromIntegral b]))

middleElement :: [Integer] -> Integer
middleElement xs = xs !! (length xs `div` 2)

isValid :: IntMap.IntMap [Integer] -> [Integer] -> Bool
isValid rules [] = True
isValid rules (x : xs) = all (\y -> notElem y $ IntMap.findWithDefault [] (fromInteger x) rules) xs && isValid rules xs

comparePages :: IntMap.IntMap [Integer] -> Integer -> Integer -> Ordering
comparePages rules a b
    | a `elem` IntMap.findWithDefault [] (fromInteger b) rules = LT
    | b `elem` IntMap.findWithDefault [] (fromInteger a) rules = GT
    | otherwise = EQ

sortInvalid :: IntMap.IntMap [Integer] -> [Integer] -> [Integer]
sortInvalid rules = sortBy (comparePages rules)

solve :: String -> Integer
solve s =
    let (rules, updates) = case parse inputParser "" s of
            Left err -> error $ show err
            Right (rules, updates) -> (rules, updates)
    in  sum $ map (middleElement . sortInvalid rules) $ filter (not . isValid rules) updates

-- since we are getting the middle element of the sorted list,
-- it doesn't matter if we sort it forewards of backwards

main :: IO ()
main = do
    s <- readFile "input.in"
    print $ solve s

CREATE TEMP TABLE input (value STRING);
INSERT INTO input VALUES (TRIM(readfile("input.in"), char(10)));

CREATE TEMP TABLE lines (s STRING);

WITH RECURSIVE
    nn (s, rest)
AS (
    SELECT
        (SELECT SUBSTR(input.value, 0, INSTR(input.value, char(10))) FROM input),
        (SELECT SUBSTR(input.value, INSTR(input.value, char(10)) + 1) FROM input)
    UNION ALL
    SELECT
        CASE INSTR(nn.rest, char(10))
            WHEN 0 THEN nn.rest
            ELSE SUBSTR(nn.rest, 0, INSTR(nn.rest, char(10)))
        END,
        CASE INSTR(nn.rest, char(10))
            WHEN 0 THEN ''
            ELSE SUBSTR(nn.rest, INSTR(nn.rest, char(10)) + 1)
        END
    FROM nn
    WHERE LENGTH(nn.rest) > 0
)
INSERT INTO lines (s)
SELECT nn.s FROM nn;
DROP TABLE input;

CREATE TABLE leftCoords (n INT);
CREATE TABLE rightCoords (n INT);

INSERT INTO leftCoords (n)
SELECT SUBSTR(s, 0, INSTR(s, " "))
FROM lines;

INSERT INTO rightCoords (n)
SELECT SUBSTR(s, INSTR(s, " "))
FROM lines;

SELECT SUM(n) AS part2
FROM (
    SELECT l.n * COUNT(r.n) AS n
    FROM leftCoords AS l
    LEFT JOIN rightCoords AS r
    ON l.n = r.n
    GROUP BY l.n
);

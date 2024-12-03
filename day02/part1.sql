CREATE TEMP TABLE input (value STRING);
INSERT INTO input VALUES (TRIM(readfile("input.in"), char(10)));

CREATE TEMP TABLE lines (id INTEGER PRIMARY KEY AUTOINCREMENT, s STRING);

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
CREATE TABLE reports (groupID INT, idx INT, num INT);

WITH RECURSIVE
    nn (groupID, idx, n, rest)
AS (
    SELECT
        lines.id AS groupID,
        0 AS idx,
        SUBSTR(lines.s, 0, INSTR(lines.s, " ")) AS n,
        SUBSTR(lines.s, INSTR(lines.s, " ") + 1) AS rest
    FROM lines
    UNION ALL
    SELECT
        nn.groupID,
        nn.idx + 1 AS idx,
        CASE INSTR(nn.rest, " ")
            WHEN 0 THEN nn.rest
            ELSE SUBSTR(nn.rest, 0, INSTR(nn.rest, " "))
        END AS n,
        CASE INSTR(nn.rest, " ")
            WHEN 0 THEN ""
            ELSE SUBSTR(nn.rest, INSTR(nn.rest, " ") + 1)
        END
    FROM nn
    WHERE nn.rest != ""
)
INSERT INTO reports (groupID, idx, num)
SELECT nn.groupID, nn.idx, nn.n
FROM nn;

SELECT COUNT(gid) AS part1
FROM (
    WITH joined AS (
        SELECT r1.groupID as gid, r1.num as n1, r2.num as n2
        FROM reports AS r1
        LEFT JOIN reports AS r2
            ON r1.groupID = r2.groupID
            AND r1.idx + 1 = r2.idx
    )
    SELECT gid
    FROM joined
    WHERE n2 IS NULL OR n1-n2 BETWEEN 1 AND 3
    GROUP BY gid
    HAVING COUNT(*) = (SELECT COUNT(*) FROM reports AS r3 WHERE r3.groupID = gid)
    UNION
    SELECT gid
    FROM joined
    WHERE n2 IS NULL OR n1-n2 BETWEEN -3 AND -1
    GROUP BY gid
    HAVING COUNT(*) = (SELECT COUNT(*) FROM reports AS r3 WHERE r3.groupID = gid)
);


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


-- SELECT * FROM reports ORDER BY groupID;

-- it is unnecessary to check if the original list is valid,
-- because if it is, then the list with its first element
-- removed is also valid.
SELECT COUNT(groupID) AS part2
FROM (
    SELECT DISTINCT r1.groupID
    FROM reports AS r1
    WHERE EXISTS (
        WITH removed_joined AS (
            SELECT
                sub.num AS n1,
                sub_next.num AS n2,
                sub.groupID AS gid
            FROM (
                SELECT
                    CASE
                        WHEN r2.idx > r1.idx THEN r2.idx - 1
                        ELSE r2.idx
                    END AS idx,
                    r2.num,
                    r2.groupID
                FROM reports r2
                WHERE r2.groupID = r1.groupID AND r2.idx != r1.idx
            ) AS sub
            LEFT JOIN (
                SELECT
                    CASE
                        WHEN r2.idx > r1.idx THEN r2.idx - 1
                        ELSE r2.idx
                    END AS idx,
                    r2.num
                FROM reports r2
                WHERE r2.groupID = r1.groupID AND r2.idx != r1.idx
            ) AS sub_next
            ON sub.idx + 1 = sub_next.idx
        )
        SELECT 1
        FROM removed_joined
        WHERE n2 IS NULL OR n1-n2 BETWEEN 1 AND 3
        GROUP BY gid
        HAVING COUNT(*) = (SELECT COUNT(*) - 1 FROM reports AS r3 WHERE r3.groupID = gid)
        UNION
        SELECT 1
        FROM removed_joined
        WHERE n2 IS NULL OR n1-n2 BETWEEN -3 AND -1
        GROUP BY gid
        HAVING COUNT(*) = (SELECT COUNT(*) - 1 FROM reports AS r3 WHERE r3.groupID = gid)
    )
);

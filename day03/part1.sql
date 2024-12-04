CREATE TEMP TABLE input (value STRING);
INSERT INTO input VALUES (TRIM(readfile("input.in"), char(10)));

CREATE TEMP TABLE muls (s STRING);
CREATE TEMP TABLE nums (n1 INTEGER, n2 INTEGER);

WITH RECURSIVE
    nn (s)
AS (
    SELECT
        (SELECT SUBSTR(input.value, INSTR(input.value, "mul(")+4) FROM input)
    UNION ALL
    SELECT
        CASE INSTR(nn.s, "mul(")
            WHEN 0 THEN NULL
            ELSE SUBSTR(nn.s, INSTR(nn.s, "mul(")+4)
        END
    FROM nn
    WHERE LENGTH(nn.s) > 0
)
INSERT INTO nums (n1, n2)
SELECT
    SUBSTR(nn.s, 0, INSTR(nn.s, ",")),
    SUBSTR(nn.s, INSTR(nn.s, ",")+1, INSTR(SUBSTR(nn.s,INSTR(nn.s, ",")+1), ")")-1)
FROM nn;

SELECT SUM(prod) AS part1
FROM (
    SELECT n1*n2 AS prod
    FROM NUMS
    WHERE
    n1 IS NOT NULL AND
    n2 IS NOT NULL AND
    LENGTH(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(n1,"0",""),"1",""),"2",""),"3",""),"4",""),"5",""),"6",""),"7",""),"8",""),"9","")) = 0 AND
    LENGTH(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(n2,"0",""),"1",""),"2",""),"3",""),"4",""),"5",""),"6",""),"7",""),"8",""),"9","")) = 0
);


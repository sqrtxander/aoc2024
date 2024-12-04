CREATE TEMP TABLE input (value STRING);
INSERT INTO input VALUES (TRIM(readfile("input.in"), char(10)));

CREATE TEMP TABLE nums (n1 INTEGER, n2 INTEGER, do INTEGER);


WITH RECURSIVE
    nn (s,do)
AS (
    SELECT
        (SELECT SUBSTR(input.value, MIN(
                    INSTR(input.value, "mul(")+4,
                    INSTR(input.value, "do()"),
                    INSTR(input.value, "don't()")
        )) FROM input),
        (SELECT TRUE)
    UNION ALL
    SELECT
        CASE INSTR(nn.s, "mul(")
            WHEN 0 THEN NULL
            ELSE SUBSTR(nn.s, MIN(
                    INSTR(nn.s, "mul(")+4,
                    CASE INSTR(nn.s, "do()")
                    WHEN 0 THEN LENGTH(nn.s)
                    WHEN 1 THEN INSTR(nn.s, "do()")+4
                    ELSE INSTR(nn.s, "do()")
                    END,
                    CASE INSTR(nn.s, "don't()")
                    WHEN 0 THEN LENGTH(nn.s)
                    WHEN 1 THEN INSTR(nn.s, "don't()")+7
                    ELSE INSTR(nn.s, "don't()")
                    END
            )
        )
        END,
        CASE INSTR(nn.s, "do()")
            WHEN 1 THEN TRUE
            ELSE CASE INSTR(nn.s, "don't()")
                WHEN 1 THEN FALSE
                ELSE nn.do
            END
        END
    FROM nn
    WHERE LENGTH(nn.s) > 0
)
INSERT INTO nums (n1, n2, do)
SELECT
    SUBSTR(nn.s, 0, INSTR(nn.s, ",")),
    SUBSTR(nn.s, INSTR(nn.s, ",")+1, INSTR(SUBSTR(nn.s,INSTR(nn.s, ",")+1), ")")-1),
    nn.do
FROM nn;

SELECT SUM(prod) AS part2
FROM (
    SELECT n1*n2 AS prod
    FROM NUMS
    WHERE
    n1 IS NOT NULL AND
    n2 IS NOT NULL AND
    LENGTH(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(n1,"0",""),"1",""),"2",""),"3",""),"4",""),"5",""),"6",""),"7",""),"8",""),"9","")) = 0 AND
    LENGTH(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(n2,"0",""),"1",""),"2",""),"3",""),"4",""),"5",""),"6",""),"7",""),"8",""),"9","")) = 0 AND
    do
);

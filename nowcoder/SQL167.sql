WITH t1 AS (SELECT DISTINCT uid,
    DATE (in_time) AS dt
   , DENSE_RANK() over(PARTITION BY uid ORDER BY DATE (in_time)) AS rn -- 编号
FROM
    tb_user_log
WHERE
    DATE (in_time) BETWEEN '2021-07-07'
  AND '2021-10-31'
  AND artical_id = 0
  AND sign_in = 1
    )
    , t2 AS (
SELECT
    *, DATE_SUB(dt, INTERVAL rn day) dt_tmp, case DENSE_RANK() OVER(PARTITION BY DATE_SUB(dt, INTERVAL rn day), uid ORDER BY dt )%7 -- 再次编号
    WHEN 3 THEN 3
    WHEN 0 THEN 7
    ELSE 1
    END AS day_coin
FROM
    t1
    )

SELECT uid,
       DATE_FORMAT(dt, '%Y%m') AS `month`,
       SUM(day_coin)           AS `coin` -- 总金币数
FROM t2
GROUP BY uid,
         DATE_FORMAT(dt, '%Y%m')
ORDER BY DATE_FORMAT(dt, '%Y%m') ASC,
         uid ASC;

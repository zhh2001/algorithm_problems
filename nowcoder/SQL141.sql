WITH t AS (SELECT
    YEAR (start_time) AS year
   , exam_id
   , tag
   , COUNT (submit_time) cnt
   , RANK() OVER(PARTITION BY YEAR (start_time) ORDER BY COUNT (submit_time) DESC) AS rk
FROM
    exam_record
    INNER JOIN
    examination_info USING (exam_id)
WHERE
    MONTH (submit_time) <= 6
GROUP BY
    year, exam_id
    )

SELECT t1.tag,
       t1.cnt                                             AS exam_cnt_20,
       t2.cnt                                             AS exam_cnt_21,
       CONCAT(ROUND((t2.cnt / t1.cnt - 1) * 100, 1), '%') AS growth_rate,
       t1.rk                                              AS exam_cnt_rank_20,
       t2.rk                                              AS exam_cnt_rank_21,
       CAST(t2.rk AS SIGNED) - CAST(t1.rk AS SIGNED)
FROM t t1
         INNER JOIN
     t t2 ON t1.year = 2020 AND t2.year = 2021 AND t1.exam_id = t2.exam_id
ORDER BY growth_rate DESC, t2.rk DESC;

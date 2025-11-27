WITH t_tag_count AS (SELECT uid,
                            `level`,
                            COUNT(start_time) - COUNT(submit_time) AS incomplete_cnt,  -- 未完成数
                            ROUND(
                                    IFNULL(1 - COUNT(submit_time) / COUNT(start_time), 0),
                                    3)                             AS incomplete_rate, -- 此人未完成率
                            COUNT(start_time)                      AS total_cnt        -- 总作答数
                     FROM exam_record
                              RIGHT OUTER JOIN user_info USING (uid)
                     GROUP BY uid)

SELECT uid, incomplete_cnt, incomplete_rate
FROM t_tag_count
WHERE EXISTS (SELECT uid
              FROM t_tag_count
              WHERE `level` = 0
                AND incomplete_cnt > 2)
  AND `level` = 0
UNION ALL
SELECT uid, incomplete_cnt, incomplete_rate
FROM t_tag_count
WHERE NOT EXISTS (SELECT uid
                  FROM t_tag_count
                  WHERE `level` = 0
                    AND incomplete_cnt > 2)
  AND total_cnt > 0
ORDER BY incomplete_rate ASC;

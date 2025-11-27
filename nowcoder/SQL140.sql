WITH incomplete AS (SELECT ui.uid,
                           SUM(
                                   CASE
                                       WHEN submit_time IS NULL THEN 1
                                       ELSE 0
                                       END
                           )                     AS incomplete_cnt,
                           COUNT(start_time)     AS total_cnt,
                           SUM(
                                   CASE
                                       WHEN submit_time IS NULL THEN 1
                                       ELSE 0
                                       END
                           ) / COUNT(start_time) AS incomplete_rate,
                           RANK()                   OVER (
                ORDER BY
                    SUM(
                        CASE
                            WHEN submit_time IS NULL THEN 1
                            ELSE 0
                        END
                    ) / COUNT(start_time) DESC
            ) i_r_rank, (SELECT COUNT(DISTINCT uid)
                         FROM user_info) AS cnt,
                           level
                    FROM exam_record AS er
                             LEFT OUTER JOIN user_info ui ON er.uid = ui.uid
                             LEFT OUTER JOIN examination_info ei ON er.exam_id = ei.exam_id
                    WHERE tag = 'SQL'
                    GROUP BY ui.uid),
     users AS (SELECT ui.uid,
                      start_time,
                      submit_time,
                      DATE_FORMAT(start_time, '%Y%m') AS start_month,
                      DENSE_RANK()                       OVER (
                PARTITION BY
                    ui.uid
                ORDER BY
                    DATE_FORMAT (start_time, '%Y%m') DESC
            ) start_month_rank
               FROM exam_record AS er
                        LEFT OUTER JOIN user_info ui ON er.uid = ui.uid
                        LEFT OUTER JOIN examination_info ei ON er.exam_id = ei.exam_id
               WHERE ui.uid IN (SELECT uid
                                FROM incomplete AS i
                                WHERE i_r_rank <= CEIL(cnt * 0.5)
                                  AND level IN (6, 7)
                                  AND i.uid = er.uid))

SELECT uid,
       start_month,
       COUNT(start_time)  AS total_cnt,
       COUNT(submit_time) AS complete_cnt
FROM users
WHERE start_month_rank <= 3
GROUP BY uid,
         start_month
ORDER BY uid ASC,
         start_month ASC;
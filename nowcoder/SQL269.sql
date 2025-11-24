SELECT l.date,
       CASE
           WHEN SUM(
                        CASE WHEN l.date < l1.date THEN 1 ELSE 0 END
                ) = 0 THEN 0.000
           ELSE ROUND(
                   SUM(
                           CASE WHEN l1.date = DATE_ADD(l.date, INTERVAL 1 DAY) THEN 1 ELSE 0 END
                   ) / SUM(
                           CASE WHEN l.date < l1.date THEN 1 ELSE 0 END
                       ), 3
                )
           END AS p
FROM login AS l
         LEFT JOIN (SELECT user_id, date
                    FROM (
                        SELECT
                        user_id, date, ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY date ASC) AS rk
                        FROM login
                        ) AS s
                    WHERE rk = 2) l1 ON l.user_id = l1.user_id
GROUP BY l.date;

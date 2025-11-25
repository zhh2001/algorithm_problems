WITH ordered_logins AS (SELECT user_id,
                               fdate,
                               ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY fdate) rn
                        FROM tb_dau
                        WHERE fdate BETWEEN '2023-01-01' AND '2023-01-31'),
     grouped_logins AS (SELECT user_id,
                               fdate,
                               rn,
                               DATE_SUB(fdate, INTERVAL rn DAY) grp
                        FROM ordered_logins)
SELECT user_id,
       MAX(consec_days) max_consec_days
FROM (SELECT user_id,
             grp,
             COUNT(*) consec_days
      FROM grouped_logins
      GROUP BY user_id, grp) consecutive
GROUP BY user_id;

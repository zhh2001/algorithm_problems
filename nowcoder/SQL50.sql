WITH login_dates AS (
    SELECT DISTINCT user_id, DATE(log_time) AS login_date
    FROM login_tb
    WHERE user_id IN (SELECT user_id FROM register_tb)
),
    numbered_logins AS (
    SELECT user_id, login_date,
    ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY login_date) AS rn
    FROM login_dates
),
    grouped_logins AS (
    SELECT user_id, COUNT(*) AS consecutive_days
    FROM numbered_logins
    GROUP BY user_id, DATE_SUB(login_date, INTERVAL rn DAY)
)

SELECT user_id
FROM grouped_logins
WHERE consecutive_days >= 3
ORDER BY user_id ASC;

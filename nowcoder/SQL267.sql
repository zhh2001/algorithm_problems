SELECT ROUND(COUNT(DISTINCT user_id) * 1.0 / (SELECT COUNT(DISTINCT user_id) FROM login), 3)
FROM login
WHERE (user_id, date)
          IN
      (SELECT user_id, DATE_ADD(MIN(date), INTERVAL 1 DAY) FROM login GROUP BY user_id);
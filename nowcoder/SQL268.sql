SELECT t.date,
       SUM(IF(t.r = 1, 1, 0))
FROM (SELECT user_id, date, ROW_NUMBER() OVER(PARTITION BY user_id ORDER BY date) AS r
      FROM login) AS t
GROUP BY date;
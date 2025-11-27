SELECT date, ROUND(AVG (if(type ="completed", 0, 1)), 3) AS p
FROM email
WHERE send_id NOT IN (SELECT id FROM user WHERE is_blacklist =1)
  AND receive_id NOT IN (SELECT id FROM user WHERE is_blacklist =1)
GROUP BY date;
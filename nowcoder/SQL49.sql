SELECT s.post,
       ROUND(AVG(TIMESTAMPDIFF(MINUTE, a.first_clockin, a.last_clockin) / 60), 3) AS work_hours
FROM staff_tb AS s
         INNER JOIN
     attendent_tb AS a ON s.staff_id = a.staff_id
WHERE a.last_clockin IS NOT NULL
  AND a.first_clockin IS NOT NULL
GROUP BY s.post
ORDER BY work_hours DESC;

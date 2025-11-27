SELECT job,
       FLOOR((COUNT(score) + 1) / 2) AS start,
       FLOOR((COUNT(score) + 2) / 2) AS end
FROM grade
GROUP BY job
ORDER BY job ASC;
SELECT b.id, b.job, b.score, b.r
FROM (SELECT job,
             (CASE WHEN cnt%2=1 THEN (cnt + 1) / 2 ELSE (cnt / 2) END) AS start,
             (CASE WHEN cnt%2=1 THEN (cnt + 1) / 2 ELSE (cnt / 2 + 1) END) AS end
      FROM (
          SELECT job, COUNT (id) AS cnt
          FROM grade
          GROUP BY job) AS d) AS a
         INNER JOIN
     (SELECT id,
             job,
             score,
             DENSE_RANK() OVER (PARTITION BY job ORDER BY score DESC) AS r
      FROM grade) AS b
     ON a.job = b.job
WHERE b.r = a.start
   OR b.r = a.end
ORDER BY b.id ASC;

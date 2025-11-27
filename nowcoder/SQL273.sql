SELECT a.id, l.name, a.score
FROM language AS l
         INNER JOIN
     (SELECT id, language_id, score, DENSE_RANK() OVER(PARTITION BY language_id ORDER BY score DESC) AS rank_num
      FROM grade) AS a
     ON l.id = a.language_id
WHERE rank_num <= 2
ORDER BY l.name ASC, a.score DESC, a.id ASC;
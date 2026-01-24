(SELECT u.name AS results
FROM Users u
LEFT JOIN MovieRating mr ON u.user_id = mr.user_id
GROUP BY u.user_id
ORDER BY COUNT(*) DESC, name ASC
LIMIT 1)

UNION ALL

(SELECT title AS results
FROM Movies m
LEFT JOIN MovieRating mr ON m.movie_id = mr.movie_id AND YEAR(mr.created_at) = 2020 AND MONTH(mr.created_at) = 2
GROUP BY mr.movie_id
ORDER BY AVG(mr.rating) DESC, title
LIMIT 1);

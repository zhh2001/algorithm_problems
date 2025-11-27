SELECT f.user_id,
       f.date,
       f.cnt
FROM (SELECT f1.user_id, f2.date, RANK() OVER(PARTITION BY user_id ORDER BY date) AS ranking, f1.cnt
      FROM (SELECT user_id, COUNT(1) AS cnt
            FROM order_info
            WHERE date > '2025-10-15' AND status = 'completed' AND product_name IN ('C++', 'Java', 'Python')
            GROUP BY user_id
            HAVING COUNT (user_id) >= 2
            ORDER BY user_id ASC) AS f1,
           order_info AS f2
      WHERE f1.user_id = f2.user_id
        AND f2.date > '2025-10-15'
        AND f2.status = 'completed'
        AND f2.product_name IN ('C++', 'Java', 'Python')) AS f
WHERE f.ranking = 1;

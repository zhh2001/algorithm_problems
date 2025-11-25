WITH sales AS (SELECT p.product_id,
                      p.name          AS product_name,
                      p.category,
                      SUM(o.quantity) AS total_sales
               FROM products AS p
                        JOIN
                    orders AS o ON p.product_id = o.product_id
               GROUP BY p.product_id, p.name, p.category)

SELECT product_name,
       total_sales,
       ROW_NUMBER() OVER (PARTITION BY category ORDER BY total_sales DESC, product_id ASC) AS category_rank
FROM sales
ORDER BY category ASC,
         total_sales DESC,
         product_id ASC;

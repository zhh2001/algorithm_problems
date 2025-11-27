SELECT product_id, CONCAT(profit_rate, "%") AS profit_rate
FROM (SELECT IFNULL(product_id, '店铺汇总')                               AS product_id,
             ROUND(100 * (1 - SUM(in_price * cnt) / SUM(price * cnt)), 1) AS profit_rate
      FROM (SELECT product_id, price, cnt, in_price
            FROM tb_order_detail
                     INNER JOIN tb_product_info USING (product_id)
                     INNER JOIN tb_order_overall USING (order_id)
            WHERE shop_id = 901 AND DATE (event_time) >= "2021-10-01") AS t_product_in_each_order
GROUP BY product_id WITH ROLLUP
HAVING profit_rate > 24.9 OR product_id IS NULL
ORDER BY product_id ASC
    ) AS t1;

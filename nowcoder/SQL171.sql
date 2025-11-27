SELECT product_id,
       ROUND(SUM(repurchase) / COUNT(repurchase), 3) AS repurchase_rate
FROM (SELECT uid, product_id, IF(COUNT(event_time) > 1, 1, 0) AS repurchase
      FROM tb_order_detail
               INNER JOIN tb_order_overall USING (order_id)
               INNER JOIN tb_product_info USING (product_id)
      WHERE tag = "零食"
        AND event_time >= (SELECT DATE_SUB(MAX(event_time), INTERVAL 89 DAY)
                           FROM tb_order_overall)
      GROUP BY uid, product_id) AS t_uid_product_info
GROUP BY product_id
ORDER BY repurchase_rate DESC,
         product_id ASC LIMIT 0, 3;

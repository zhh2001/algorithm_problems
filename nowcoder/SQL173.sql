SELECT dt, ROUND(cnt / total_cnt, 3) AS sale_rate, ROUND(1 - cnt / total_cnt, 3) AS unsale_rate
FROM (SELECT DISTINCT
          DATE (event_time) AS dt,
     (SELECT COUNT(DISTINCT (IF(shop_id != 901, NULL, product_id)))
      FROM tb_order_overall
               INNER JOIN tb_order_detail USING (order_id)
               INNER JOIN tb_product_info USING (product_id)
      WHERE TIMESTAMPDIFF(DAY, event_time, to1.event_time) BETWEEN 0 AND 6) AS cnt,
     (SELECT COUNT(DISTINCT product_id)
      FROM tb_product_info
      WHERE shop_id = 901) AS total_cnt FROM tb_order_overall to1
WHERE DATE (event_time) BETWEEN '2021-10-01'
  AND '2021-10-03'
    ) AS t0
ORDER BY dt ASC;

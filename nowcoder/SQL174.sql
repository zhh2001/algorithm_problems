SELECT "北京"                   AS city,
       ROUND(AVG(order_num), 3) AS avg_order_num,
       ROUND(AVG(income), 3)    AS avg_income
FROM (SELECT driver_id, COUNT(order_id) AS order_num, SUM(fare) AS income
      FROM tb_get_car_order
               INNER JOIN tb_get_car_record USING (order_id)
      WHERE city = "北京"
        and DATE_FORMAT(order_time, "%Y%m%d") BETWEEN '20211001' AND '20211007'
      GROUP BY driver_id
      HAVING COUNT(order_id) >= 3) AS t_driver_info;

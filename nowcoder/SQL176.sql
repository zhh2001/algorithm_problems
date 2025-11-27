SELECT city,
       driver_id,
       avg_grade,
       avg_order_num,
       avg_mileage
FROM (SELECT city,
             driver_id,
             ROUND(avg_grade, 1)                AS avg_grade,
             ROUND(order_num / work_days, 1)    AS avg_order_num,
             ROUND(toal_mileage / work_days, 3) AS avg_mileage,
             RANK()                                over(PARTITION BY city ORDER BY avg_grade DESC) AS rk
      FROM (SELECT driver_id,
                   city,
                   AVG(grade)                        AS avg_grade,
                   COUNT(DISTINCT DATE (order_time)) AS work_days,
                   COUNT(order_time)                 AS order_num,
                   SUM(mileage)                      AS toal_mileage
            FROM tb_get_car_record
                     INNER JOIN tb_get_car_order USING (order_id)
            GROUP BY driver_id, city) AS t_driver_info) AS t_driver_rk
WHERE rk = 1
ORDER BY avg_order_num ASC;

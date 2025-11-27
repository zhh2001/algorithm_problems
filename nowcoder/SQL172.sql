SELECT ROUND(AVG(total_amount), 1)              AS avg_amount,
       ROUND(AVG(raw_amount - total_amount), 1) AS avg_cost
FROM (SELECT uid, total_amount, raw_amount
      FROM (SELECT DISTINCT uid,
                            FIRST_VALUE(event_time) over(wd_uid_first) AS event_time, FIRST_VALUE(order_id) over(wd_uid_first) AS order_id, FIRST_VALUE(total_amount) over(wd_uid_first) AS total_amount
            FROM tb_order_overall WINDOW wd_uid_first AS (partition by uid ORDER BY event_time ASC)) AS t_first_order
               INNER JOIN (SELECT order_id, SUM(price * cnt) AS raw_amount
                           FROM tb_order_detail
                           GROUP BY order_id) AS t_raw_amount
                          USING (order_id)
      WHERE DATE_FORMAT(event_time, "%Y-%m") = '2021-10') AS t_first_order_info;

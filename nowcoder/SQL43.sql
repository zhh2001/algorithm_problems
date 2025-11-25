SELECT DATE_FORMAT(t.t_time, '%Y-%m') `time`,
       ROUND(SUM(t.t_amount), 1)      total
FROM trade t
         JOIN
     customer c ON t.t_cus = c.c_id
WHERE t.t_type = 1
          AND c.c_name = 'TOM'
          AND YEAR (t.t_time) = 2023
GROUP BY
    `time`
ORDER BY
    `time` ASC;

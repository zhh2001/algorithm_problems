SELECT start_month,
       count(distinct uid) AS mau,
       sum(new)            AS month_add_uv,
       max(sum(new))          over(order by start_month) AS max_monthe_add_uv , sum(sum(new)) over(order by start_month)AS cum_sum_uv
FROM (SELECT *
           , date_format(start_time, '%Y%m')                              AS start_month
           , if(start_time = min(start_time)over(partition by uid), 1, 0) AS new
      FROM exam_record) AS a
GROUP BY start_month
ORDER BY start_month ASC;

SELECT r1.job, r1.first_year_mon, r1.first_year_cnt, r2.second_year_mon, r2.second_year_cnt
FROM (SELECT job,
             DATE_FORMAT(date, "%Y-%m") AS first_year_mon,
             SUM(num)                   AS first_year_cnt
      FROM resume_info
      WHERE year (date) = 2025
      GROUP BY job, first_year_mon) AS r1
         JOIN (SELECT job,
                      DATE_FORMAT(date, "%Y-%m") AS second_year_mon,
                      SUM(num)                   AS second_year_cnt
               FROM resume_info
               WHERE year (date) = 2026
               GROUP BY job, second_year_mon) r2
              ON right (r1.first_year_mon, 2) = right (r2.second_year_mon, 2) AND r1.job = r2.job
ORDER BY r1.first_year_mon DESC, r1.job DESC;

SELECT s1.emp_no,
       s2.salary - s1.salary AS growth
FROM (SELECT e.emp_no,
             s.salary
      FROM employees e
               LEFT OUTER JOIN salaries s
                               ON e.emp_no = s.emp_no
      WHERE e.hire_date = s.from_date) AS s1 -- 入职薪水
         INNER JOIN
     (SELECT e.emp_no,
             s.salary
      FROM employees AS e
               LEFT JOIN salaries AS s
                         ON e.emp_no = s.emp_no
      WHERE s.to_date = '9999-01-01') AS s2 -- 现在薪水
     ON s1.emp_no = s2.emp_no
ORDER BY growth ASC;
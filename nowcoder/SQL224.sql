SELECT a.emp_no, b.manager_no, a.emp_salary, b.manager_salary
FROM (SELECT de.dept_no, de.emp_no, s.salary AS emp_salary
      FROM dept_emp AS de,
           salaries AS s
      WHERE de.emp_no = s.emp_no
        AND s.to_date = '9999-01-01') AS a,
     (SELECT dm.dept_no, dm.emp_no AS manager_no, s.salary AS manager_salary
      FROM dept_manager AS dm,
           salaries AS s
      WHERE dm.emp_no = s.emp_no
        AND s.to_date = '9999-01-01') AS b
WHERE a.dept_no = b.dept_no
  AND a.emp_salary > b.manager_salary;
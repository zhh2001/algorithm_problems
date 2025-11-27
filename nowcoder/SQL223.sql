SELECT b.dept_no,
       a.emp_no,
       d.salary
FROM employees AS a,
     dept_emp AS b,
     dept_manager AS c,
     salaries AS d
WHERE a.emp_no = b.emp_no
  AND b.dept_no = c.dept_no
  AND b.emp_no != c.emp_no
    AND a.emp_no = d.emp_no;
SELECT e.emp_no,
       s.salary,
       e.last_name,
       e.first_name
FROM employees AS e
         LEFT JOIN salaries s ON e.emp_no = s.emp_no
WHERE s.salary = (
SELECT MAX(salary)
    FROM salaries
    WHERE salary != (SELECT MAX (salary) FROM salaries)
);

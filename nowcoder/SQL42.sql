SELECT ct.pay_ability,
       CONCAT(ROUND(SUM(CASE WHEN lt.overdue_days IS NOT NULL THEN 1 ELSE 0 END) / COUNT(*) * 100, 1),
              '%') overdue_ratio
FROM loan_tb lt
         JOIN
     customer_tb ct ON lt.customer_id = ct.customer_id
GROUP BY ct.pay_ability
ORDER BY overdue_ratio DESC;

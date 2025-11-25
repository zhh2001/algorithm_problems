SELECT SUM(`num`) AS `staff_nums`
FROM (SELECT CASE
                 WHEN course IS NULL
                     THEN 0
                 ELSE LENGTH(course) - LENGTH(REPLACE(course, ',', '')) + 1
                 END AS `num`
      FROM `cultivate_tb`) AS `sub`;
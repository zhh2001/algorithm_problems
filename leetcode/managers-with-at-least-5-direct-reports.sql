SELECT `manager`.`name` AS `name`
FROM Employee AS `manager` INNER JOIN Employee AS `report`
ON `manager`.`id` = `report`.`managerId`
GROUP BY `manager`.`id`
HAVING COUNT(`report`.`id`) >= 5;
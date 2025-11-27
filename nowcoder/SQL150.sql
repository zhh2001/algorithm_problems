WITH t1 AS
         (SELECT level,
                 score,
                 COUNT(level) OVER(PARTITION BY level) AS level_cn, (CASE
                                                                         WHEN score < 60 THEN '差'
                                                                         WHEN score < 75 and score >= 60 THEN '中'
                                                                         WHEN score < 90 and score >= 75 THEN '良'
                                                                         ELSE '优' END) AS grade -- 对用户的分数进行等级定义
          FROM exam_record AS a
                   LEFT OUTER JOIN user_info b ON a.uid = b.uid
          WHERE score IS NOT NULL -- 剔除未完成的试卷
         )

SELECT level, grade, ROUND(COUNT(grade) / level_cn, 3) AS ratio
FROM t1
GROUP BY level, grade
ORDER BY level DESC, ratio DESC;
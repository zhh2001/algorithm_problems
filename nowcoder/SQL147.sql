SELECT uid,
       nick_name,
       achievement
FROM user_info
WHERE nick_name LIKE '牛客%号'
  AND achievement BETWEEN 1200 AND 2500
  AND uid IN (SELECT uid
              FROM (SELECT uid, start_time AS act_time
                    FROM exam_record
                    UNION ALL
                    SELECT uid, submit_time AS act_time
                    FROM practice_record) AS temp
              GROUP BY uid
              HAVING DATE_FORMAT(MAX(act_time), '%Y%m') = 202109);

SELECT uid, `level`, `register_time`, max_score
FROM user_info
         INNER JOIN exam_record USING (uid)
         INNER JOIN examination_info USING (exam_id)
         INNER JOIN (SELECT uid, MAX(score) AS max_score
                     FROM exam_record
                     GROUP BY uid) AS t_exam_max_score USING (uid)
WHERE `job` = '算法'
  AND `tag` = '算法'
  AND DATE (`register_time`)= DATE (submit_time)
ORDER BY max_score DESC
    LIMIT 6, 3;

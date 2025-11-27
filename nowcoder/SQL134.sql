SELECT uid, exam_cnt, IFNULL(question_cnt, 0) AS question_cnt
FROM (SELECT uid
      FROM exam_record
               INNER JOIN examination_info USING (exam_id)
               INNER JOIN user_info USING (uid)
      WHERE difficulty = 'hard'
        AND tag = 'SQL'
        AND `level` = 7
      GROUP BY uid
      HAVING AVG(score) > 80) AS t_user_id
         INNER JOIN (SELECT uid, COUNT(exam_id) AS exam_cnt
                     FROM exam_record
                     WHERE YEAR (submit_time)=2021 AND submit_time IS NOT NULL
                     GROUP BY uid) AS t_exam_cnt
                    USING (uid)
         LEFT OUTER JOIN (SELECT uid, COUNT(question_id) AS question_cnt
                          FROM practice_record
                          WHERE YEAR (submit_time)=2021
                          GROUP BY uid) AS t_question_cnt USING (uid)
ORDER BY exam_cnt ASC, question_cnt DESC;

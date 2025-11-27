SELECT MAX(tag)        AS tag,
       MAX(answer_cnt) AS answer_cnt
FROM (SELECT MAX(er.exam_id) AS exam_id,
             ei.tag          AS tag,
             COUNT(1)        AS answer_cnt
      FROM exam_record AS er
               INNER JOIN examination_info ei ON er.exam_id = ei.exam_id
      GROUP BY tag) t
GROUP BY UPPER(tag)
HAVING COUNT(exam_id) > 1
ORDER BY answer_cnt DESC;

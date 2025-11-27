SELECT uid, exam_id, ROUND(AVG(score), 0) avg_score
FROM exam_record
WHERE uid IN (SELECT uid
              FROM user_info
              WHERE nick_name
    RLIKE "^牛客[0-9]+号$"
   OR nick_name RLIKE "^[0-9]+$")
  AND
    exam_id IN (
SELECT exam_id
FROM examination_info
WHERE tag RLIKE "^[cC]")
  AND
    score IS NOT NULL
GROUP BY uid, exam_id
ORDER BY uid, avg_score ASC;

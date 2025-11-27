SELECT uid, COUNT(start_time) AS exam_complete_cnt
FROM (SELECT uid, start_time, submit_time
      FROM (SELECT uid,
                   start_time,
                   submit_time,
                   DENSE_RANK() OVER (PARTITION BY uid ORDER BY DATE_FORMAT(start_time,'%Y-%m') DESC) AS ranking
            FROM exam_record) AS a1
      WHERE ranking <= 3) AS a2
GROUP BY uid
HAVING COUNT(start_time) = COUNT(submit_time) -- 必须相等才能说明没有试卷是未完成状态
ORDER BY exam_complete_cnt DESC, uid DESC;

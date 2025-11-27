SELECT uid, days_window, ROUND(days_window * exam_cnt / diff_days, 2) AS avg_exam_cnt
FROM (SELECT uid,
             COUNT(start_time)                              AS exam_cnt,   -- 此人作答的总试卷数
             DATEDIFF(MAX(start_time), MIN(start_time)) + 1 AS diff_days,  -- 最早一次作答和最晚一次作答的相差天数
             MAX(DATEDIFF(next_start_time, start_time)) + 1 AS days_window -- 两次作答的最大时间窗
      FROM (SELECT uid,
                   exam_id,
                   start_time,
                   LEAD(start_time) OVER(
                PARTITION BY uid ORDER BY start_time) AS next_start_time -- 将连续的下次作答时间拼上
            FROM exam_record
            WHERE YEAR (start_time)=2021) AS t_exam_record_lead
      GROUP BY uid) AS t_exam_record_stat
WHERE diff_days > 1
ORDER BY days_window DESC, avg_exam_cnt DESC;

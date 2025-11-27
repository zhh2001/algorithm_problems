SELECT user_info.uid                     AS uid,
       IFNULL(act_month_total, 0)        AS act_month_total,
       IFNULL(act_days_2021, 0)          AS act_days_2021,
       IFNULL(act_days_2021_exam, 0)     AS act_days_2021_exam,
       IFNULL(act_days_2021_question, 0) AS act_days_2021_question
FROM user_info
         LEFT OUTER JOIN (SELECT uid,
                                 COUNT(DISTINCT DATE_FORMAT(active_date, "%Y%m"))              AS act_month_total,
                                 COUNT(DISTINCT IF(YEAR(active_date)=2021, active_date, NULL)) AS act_days_2021
                          FROM (SELECT DISTINCT uid, DATE (start_time) AS active_date FROM exam_record
                          UNION ALL
                          SELECT DISTINCT uid, DATE (submit_time) AS active_date
                          FROM practice_record) AS t_merge_record
GROUP BY uid ) AS t_2021_total_act USING(uid) -- 总活跃月份数、2021年活跃天数
LEFT OUTER JOIN (
    SELECT uid, COUNT(DISTINCT DATE(submit_time)) AS act_days_2021_question
    FROM practice_record
    WHERE YEAR(submit_time)=2021
    GROUP BY uid
) AS t_2021_act_days_question USING(uid) -- 2021年题目练习活跃天数
LEFT OUTER JOIN (
    SELECT uid, COUNT(DISTINCT DATE(start_time)) AS act_days_2021_exam
    FROM exam_record
    WHERE YEAR(start_time)=2021
    GROUP BY uid
) AS t_2021_act_days_exam USING(uid) -- 2021试卷作答活跃天数
WHERE user_info.level > 5
ORDER BY act_month_total DESC, act_days_2021 DESC;

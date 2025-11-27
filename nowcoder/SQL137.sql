SELECT exam_id, duration, release_time
FROM (SELECT DISTINCT exam_id,
                      duration,
                      release_time,
                      NTH_VALUE(time_took, 2) OVER (PARTITION BY exam_id ORDER BY time_took DESC) AS max2_time_took,
                      NTH_VALUE(time_took, 2) OVER (PARTITION BY exam_id ORDER BY time_took ASC)  AS min2_time_took
      FROM (SELECT uid,
                   exam_record.exam_id                                 AS exam_id,
                   start_time,
                   duration,
                   release_time,
                   TimeStampDiff(SECOND, start_time, submit_time) / 60 AS time_took
            FROM exam_record
                     LEFT OUTER JOIN examination_info USING (exam_id)
            WHERE submit_time IS NOT NULL) AS t_exam_record_timetook) AS t_exam_time_took
WHERE max2_time_took IS NOT NULL
  AND min2_time_took IS NOT NULL
  AND max2_time_took - min2_time_took > duration / 2
ORDER BY exam_id DESC;

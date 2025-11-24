SELECT video_id,
       ROUND((100 * finish_rate + 5 * like_index + 3 * comment_index + 2 * retweet_index) / (fresh_index + 1),
             0) hot_index
FROM (SELECT a.video_id,
             SUM(IF(timestampdiff(second,a.start_time, a.end_time) - b.duration >= 0, 1, 0)) /
             COUNT(a.video_id)                                                                        AS finish_rate,
             SUM(a.if_like)                                                                           AS like_index,
             SUM(IF(a.comment_id is not null, 1, 0))                                                  AS comment_index,
             SUM(a.if_retweet)                                                                        AS retweet_index,
             IF(COUNT(a.video_id) = 0,
                DATEDIFF(date((SELECT MAX(end_time) FROM tb_user_video_log)),date(b.release_time)),
                DATEDIFF(date((SELECT MAX(end_time) FROM tb_user_video_log)),MAX(date (a.end_time)))) AS fresh_index
      FROM tb_user_video_log a
               LEFT JOIN tb_video_info b
                         ON a.video_id = b.video_id
      where DATEDIFF(date((select max(end_time) from tb_user_video_log)),DATE(b.release_time)) <= 29
      GROUP BY a.video_id) AS fir_sheet
ORDER BY hot_index DESC limit 0,3

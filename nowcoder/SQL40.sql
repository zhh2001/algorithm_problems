SELECT sub.month,
       sub.ranking,
       sub.song_name,
       sub.play_pv
FROM (SELECT
          MONTH (p.fdate) month, ROW_NUMBER() OVER (
          PARTITION BY
          MONTH (p.fdate)
          ORDER BY
          COUNT (*) DESC, s.song_id
          ) ranking, s.song_name, COUNT (*) play_pv
      FROM
          play_log p
          JOIN song_info s
      ON s.song_id = p.song_id
          JOIN user_info u ON u.user_id = p.user_id
      WHERE
          s.singer_name = '周杰伦'
        AND u.age BETWEEN 18
        AND 25
        AND YEAR (p.fdate) = 2022
      GROUP BY
          MONTH (p.fdate),
          s.song_name,
          s.song_id) sub
WHERE sub.ranking <= 3
ORDER BY sub.month,
         sub.play_pv DESC;

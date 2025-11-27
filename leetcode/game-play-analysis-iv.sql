SELECT IFNULL(ROUND(COUNT(DISTINCT (Result.player_id)) / COUNT(DISTINCT (Activity.player_id)), 2), 0) AS fraction
FROM (SELECT Activity.player_id AS player_id
      FROM (SELECT player_id, DATE_ADD(MIN(event_date), INTERVAL 1 DAY) AS second_date
            FROM Activity
            GROUP BY player_id) AS Expected,
           Activity
      WHERE Activity.event_date = Expected.second_date
        AND Activity.player_id = Expected.player_id) AS Result,
     Activity;
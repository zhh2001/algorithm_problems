SELECT ct.user_id,
       ct.room_id,
       gt.room_type,
       DATEDIFF(ct.checkout_time, ct.checkin_time) days
FROM checkin_tb ct
         INNER JOIN
     guestroom_tb gt ON gt.room_id = ct.room_id
WHERE DATEDIFF(ct.checkout_time, ct.checkin_time) >= 2
  AND DATE (ct.checkin_time) = '2022-06-12'
ORDER BY
    days, ct.room_id, ct.user_id DESC;

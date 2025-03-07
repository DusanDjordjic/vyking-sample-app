SELECT
    id,
    name,
    prize,
    start_date,
    end_date
FROM
    tournaments
WHERE
    id = ?;

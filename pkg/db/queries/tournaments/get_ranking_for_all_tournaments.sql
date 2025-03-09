WITH ranked_results AS (
	SELECT
		player_id,
		tournament_id,
		RANK() OVER (PARTITION BY tournament_id ORDER BY SUM(bet_amount) DESC) as place
	FROM tournament_bets bets
	GROUP BY player_id, tournament_id
),
winnings_results AS (SELECT
		player_id,
		tournament_id,
		prize,
		place,
		CASE
			WHEN place = 1 THEN prize * 0.5
			WHEN place = 2 THEN prize * 0.3
			WHEN place = 3 THEN prize * 0.2
			ELSE 0
		END AS winnings
	FROM ranked_results res
	JOIN tournaments ts ON ts.id = res.tournament_id
)
SELECT
	player_id,
	SUM(winnings) as total_winnings
FROM winnings_results
GROUP BY player_id
ORDER BY total_winnings DESC;

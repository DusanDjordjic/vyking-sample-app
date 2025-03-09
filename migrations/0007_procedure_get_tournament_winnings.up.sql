CREATE PROCEDURE GetTournamentWinnings(IN tid BIGINT)
BEGIN
	-- create a temp table to store winnigs
	CREATE TEMPORARY TABLE temp_winnigs (
		player_id BIGINT,
		tournament_id BIGINT,
		total_bets DOUBLE,
		winnings DOUBLE
	);

	CALL CalcTournamentWinnings(tid);

	SELECT player_id, winnings FROM temp_winnigs;

	DROP TEMPORARY TABLE temp_winnigs;
END;

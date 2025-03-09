CREATE PROCEDURE GetAllTournamentWinnings()
BEGIN

	DECLARE current_tid BIGINT;
	DECLARE done INT DEFAULT 0;

	DECLARE tournament_cursor CURSOR FOR
		SELECT id FROM tournaments
		WHERE start_date < CURRENT_TIMESTAMP;

	DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = 1;

	OPEN tournament_cursor;

	-- create a temp table to store winnigs
	CREATE TEMPORARY TABLE temp_winnigs (
		player_id BIGINT,
		tournament_id BIGINT,
		total_bets DOUBLE,
		winnings DOUBLE
	);

	tournament_loop: LOOP
		FETCH tournament_cursor INTO current_tid;

		IF done = 1 THEN
			LEAVE tournament_loop;
		END IF;

		CALL CalcTournamentWinnings(current_tid);

	END LOOP;


	CLOSE tournament_cursor;

	SELECT player_id, SUM(winnings) as total_winnings FROM temp_winnigs GROUP BY player_id ORDER BY total_winnings DESC;

	DROP TEMPORARY TABLE temp_winnigs;
END;

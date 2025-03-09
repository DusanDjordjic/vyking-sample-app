CREATE PROCEDURE CalculateTurnamentWinnings(IN tid BIGINT)
BEGIN
	DECLARE total_prize DOUBLE DEFAULT 0;
	DECLARE first_prize DOUBLE DEFAULT 0;
	DECLARE second_prize DOUBLE DEFAULT 0;
	DECLARE third_prize DOUBLE DEFAULT 0;

	-- variable used to store cursor values
	DECLARE current_pid BIGINT;
	DECLARE current_pbet DOUBLE;

	-- players to loop through
	DECLARE player_cursor CURSOR FOR
		SELECT player_id, SUM(bet_amount) as total_bets
		FROM tournament_bets
		WHERE tournament_id = tid
		GROUP BY player_id
		ORDER BY total_bets DESC LIMIT 3;

	-- get the total prize
	SELECT prize INTO total_prize FROM tournaments
	WHERE id = tid;

	-- calc the winnings
	SET first_prize = total_prize * 0.5;
	SET second_prize = total_prize * 0.3;
	SET third_prize = total_prize * 0.2;

	-- create a temp table to store winnigs
	CREATE TEMPORARY TABLE temp_winnigs (
		player_id BIGINT,
		total_bets DOUBLE,
		winnings DOUBLE
	);

	OPEN player_cursor;

	FETCH player_cursor INTO current_pid, current_pbet;
	INSERT INTO temp_winnigs (player_id, winnings, total_bets) VALUES (current_pid, first_prize, current_pbet);

	FETCH player_cursor INTO current_pid, current_pbet;
	INSERT INTO temp_winnigs (player_id, winnings, total_bets) VALUES (current_pid, second_prize, current_pbet);

	FETCH player_cursor INTO current_pid, current_pbet;
	INSERT INTO temp_winnigs (player_id, winnings, total_bets) VALUES (current_pid, third_prize, current_pbet);

	CLOSE player_cursor;

	SELECT player_id, winnings, total_bets FROM temp_winnigs;
	DROP TEMPORARY TABLE temp_winnigs;
END;

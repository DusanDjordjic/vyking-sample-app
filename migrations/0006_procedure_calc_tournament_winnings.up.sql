CREATE PROCEDURE CalcTurnamentEarnings(IN tid BIGINT)
BEGIN
	DECLARE current_rank INT DEFAULT 1;
	DECLARE total_prize DOUBLE DEFAULT 0;
	DECLARE first_prize DOUBLE DEFAULT 0;
	DECLARE second_prize DOUBLE DEFAULT 0;
	DECLARE third_prize DOUBLE DEFAULT 0;

	-- variable used to store cursor values
	DECLARE current_pid BIGINT;
	DECLARE current_pbet DOUBLE;

	DECLARE done INT DEFAULT 0;

	-- players to loop through
	DECLARE player_cursor CURSOR FOR
		SELECT player_id, SUM(bet_amount) as total_bets
		FROM tournament_bets
		WHERE tournament_id = tid
		GROUP BY player_id
		ORDER BY total_bets DESC LIMIT 3;

	DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = 1;

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
		winnings DOUBLE
	);

	OPEN player_cursor;

	rank_loop: LOOP
		FETCH player_cursor INTO current_pid, current_pbet;

		IF done = 1 THEN
			LEAVE rank_loop;
		END IF;

		IF current_rank > 3 THEN
			LEAVE rank_loop;
		END IF;

		IF  current_rank = 1 THEN
			INSERT INTO temp_winnigs (player_id, winnings) VALUES (current_pid, first_prize);
		ELSEIF current_rank = 2 THEN
			INSERT INTO temp_winnigs (player_id, winnings) VALUES (current_pid, second_prize);
		ELSE
			INSERT INTO temp_winnigs (player_id, winnings) VALUES (current_pid, third_prize);
		END IF;

		-- update the rank
		SET current_rank = current_rank + 1;

	END LOOP;


	CLOSE player_cursor;

	SELECT player_id, winnings FROM temp_winnigs ORDER BY winnings DESC;
	DROP TEMPORARY TABLE temp_winnigs;
END;

DROP procedure CalcTurnamentEarnings;

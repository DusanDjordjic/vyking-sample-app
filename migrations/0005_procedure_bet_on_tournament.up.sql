CREATE PROCEDURE BetOnTournament(IN player_id BIGINT, IN tournament_id BIGINT, IN amount DOUBLE)
BEGIN
	DECLARE player_account_balance DOUBLE DEFAULT 0;
	DECLARE last_id BIGINT DEFAULT 0;

	DECLARE t_start_date DATETIME;
	DECLARE t_end_date DATETIME;

	-- If we couldn't get the tournament by id
	DECLARE EXIT HANDLER FOR NOT FOUND
	BEGIN
		ROLLBACK;
		SIGNAL SQLSTATE VALUE '45002' SET MESSAGE_TEXT = "tournament doesn't exist";
	END;

	-- if user doesn't have enough funds
	DECLARE EXIT HANDLER FOR SQLSTATE VALUE '45001'
	BEGIN
		ROLLBACK;
		RESIGNAL SET MESSAGE_TEXT = "insufficient funds";
	END;

	-- if tournament hasn't started yet
	DECLARE EXIT HANDLER FOR SQLSTATE VALUE '45003'
	BEGIN
		ROLLBACK;
		RESIGNAL SET MESSAGE_TEXT = "tournament hasn't started yet";
	END;

	-- if tournament already ended
	DECLARE EXIT HANDLER FOR SQLSTATE VALUE '45004'
	BEGIN
		ROLLBACK;
		RESIGNAL SET MESSAGE_TEXT = "tournament has already ended";
	END;

	-- other error, never called probably
	DECLARE EXIT HANDLER FOR SQLEXCEPTION
	BEGIN
		ROLLBACK;
		RESIGNAL;
	END;

	START TRANSACTION;

	SELECT account_balance INTO player_account_balance
	FROM players
	WHERE id = player_id FOR UPDATE;

	IF player_account_balance < amount THEN
		ROLLBACK;
		SIGNAL SQLSTATE '45001';
	ELSEIF player_account_balance = 0 THEN
		ROLLBACK;
		SIGNAL SQLSTATE '45001';
	END IF;


	-- check if turnament is started or not
	SELECT start_date, end_date INTO t_start_date, t_end_date FROM tournaments
	WHERE id = tournament_id;


	SELECT t_start_date, t_end_date;

	IF t_start_date > CURRENT_TIMESTAMP THEN
		ROLLBACK;
		SIGNAL SQLSTATE '45003';
	END IF;

	IF t_end_date <= CURRENT_TIMESTAMP THEN
		ROLLBACK;
		SIGNAL SQLSTATE '45004';
	END IF;

	UPDATE players SET account_balance = account_balance - amount WHERE id = player_id;

	INSERT INTO tournament_bets (player_id, tournament_id, bet_amount) VALUES (player_id, tournament_id, amount);

	SET last_id = LAST_INSERT_ID();

	SELECT id, created_at, player_id, tournament_id, bet_amount
	FROM tournament_bets
	WHERE id = last_id;

	COMMIT;
END;

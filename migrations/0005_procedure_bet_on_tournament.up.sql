CREATE PROCEDURE BetOnTournament(IN player_id BIGINT, IN tournament_id BIGINT, IN amount DOUBLE)
BEGIN
	DECLARE player_account_balance DOUBLE DEFAULT 0;
	DECLARE last_id BIGINT DEFAULT 0;

	DECLARE EXIT HANDLER FOR SQLSTATE VALUE '45001'
	BEGIN
		ROLLBACK;
		RESIGNAL SET MYSQL_ERRNO = 10000, MESSAGE_TEXT = "insufficient funds";
	END;

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

	UPDATE players SET account_balance = account_balance - amount WHERE id = player_id;

	INSERT INTO tournament_bets (player_id, tournament_id, bet_amount) VALUES (player_id, tournament_id, amount);

	SET last_id = LAST_INSERT_ID();

	SELECT id, created_at, player_id, tournament_id, bet_amount
	FROM tournament_bets
	WHERE id = last_id;

	COMMIT;
END;

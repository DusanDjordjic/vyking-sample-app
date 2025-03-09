CREATE TABLE IF NOT EXISTS tournament_bets (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    tournament_id BIGINT NOT NULL,
    player_id BIGINT NOT NULL,
    bet_amount DOUBLE NOT NULL,
    CONSTRAINT positive_bet_amount CHECK (bet_amount > 0),
    CONSTRAINT fk_tournaments FOREIGN KEY (tournament_id) REFERENCES tournaments (id),
    CONSTRAINT fk_players FOREIGN KEY (player_id) REFERENCES players (id)
);

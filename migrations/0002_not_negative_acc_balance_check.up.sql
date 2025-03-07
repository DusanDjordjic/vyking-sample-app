ALTER TABLE players ADD CONSTRAINT non_negative_acc_balance CHECK (account_balance >= 0);

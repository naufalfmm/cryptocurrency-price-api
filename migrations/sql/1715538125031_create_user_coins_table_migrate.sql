CREATE TABLE "user_coins" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "user_id" INTEGER NOT NULL,
    "coin_id" INTEGER NOT NULL,
    "created_at" DATETIME NOT NULL DEFAULT current_timestamp,
	"updated_at" DATETIME NOT NULL DEFAULT current_timestamp,
	"deleted_at" DATETIME NULL,
	"created_by" TEXT NOT NULL,
	"updated_by" TEXT NOT NULL,
	"deleted_by" TEXT NULL,
	"deleted_unix" INTEGER NOT NULL DEFAULT 0,
	
	CONSTRAINT "uq_user_coins" UNIQUE(user_id, coin_id, deleted_unix),
	CONSTRAINT "fk_user_coin_user" FOREIGN KEY (user_id) REFERENCES users(id),
	CONSTRAINT "fk_user_coin_coin" FOREIGN KEY (coin_id) REFERENCES coins(id)
);

CREATE INDEX "idx_user_coin_user" ON user_coins(user_id);
CREATE INDEX "idx_user_coin_coin" ON user_coins(coin_id);
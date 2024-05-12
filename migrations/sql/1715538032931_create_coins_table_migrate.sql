CREATE TABLE "coins" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "code" TEXT NOT NULL,
    "coincap_id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "latest_price" NUMERIC(32,16) NOT NULL,
    "created_at" DATETIME NOT NULL DEFAULT current_timestamp,
	"updated_at" DATETIME NOT NULL DEFAULT current_timestamp,
	"created_by" TEXT NOT NULL,
	"updated_by" TEXT NOT NULL,
	
	CONSTRAINT "uq_coin_code" UNIQUE(code),
	CONSTRAINT "uq_coin_coincap_id" UNIQUE(coincap_id)   
);
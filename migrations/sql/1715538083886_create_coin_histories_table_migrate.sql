CREATE TABLE "coin_histories" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "coin_id" INTEGER NOT NULL,
    "latest_price" NUMERIC(32,16) NOT NULL,
    "created_at" DATETIME NOT NULL DEFAULT current_timestamp,
	"created_by" TEXT NOT NULL
);
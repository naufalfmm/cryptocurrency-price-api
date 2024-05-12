CREATE TABLE "users" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "created_at" DATETIME NOT NULL DEFAULT current_timestamp,
	"updated_at" DATETIME NOT NULL DEFAULT current_timestamp,
	"created_by" TEXT NOT NULL,
	"updated_by" TEXT NOT NULL,

    CONSTRAINT "uq_user_email" UNIQUE (email)
);
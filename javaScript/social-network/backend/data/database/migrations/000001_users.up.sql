CREATE TABLE IF NOT EXISTS "users" (
    "user_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "date_of_birth" VARCHAR(255) NOT NULL,
    "avatar" VARCHAR(255) NOT NULL DEFAULT "avatar.png",
    "nickname" VARCHAR(255),
    "about" TEXT,
    "public" INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS "sessions" (
    "session_id" VARCHAR(255) PRIMARY KEY,
    "user_id"   INTEGER NOT NULL,
    "timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
);

CREATE TABLE IF NOT EXISTS "follows" (
    "user_id" INTEGER NOT NULL,
    "follow_id" INTEGER NOT NULL,
    "confirmed" INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id"),
    FOREIGN KEY ("follow_id") REFERENCES "users" ("user_id")
);
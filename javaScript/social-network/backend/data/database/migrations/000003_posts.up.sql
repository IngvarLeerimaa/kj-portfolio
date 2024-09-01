CREATE TABLE IF NOT EXISTS "posts" (
    "post_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "user_id" INTEGER NOT NULL,
    "privacy" INTEGER NOT NULL,
    "text" TEXT NOT NULL,
    "image" VARCHAR(255) NOT NULL,
    "timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
);

CREATE TABLE IF NOT EXISTS "comments" (
    "comment_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "post_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "text" TEXT NOT NULL,
    "image" VARCHAR(255) NOT NULL,
    "timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
);

CREATE TABLE IF NOT EXISTS "specific_followers" (
    "post_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
);
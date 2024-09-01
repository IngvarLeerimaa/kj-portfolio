CREATE TABLE IF NOT EXISTS "chat" (
    "to_id" INTEGER NOT NULL,
    "from_id" INTEGER NOT NULL,
    "message" TEXT NOT NULL,
    "timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("to_id") REFERENCES "users" ("user_id"),
    FOREIGN KEY ("from_id") REFERENCES "users" ("user_id")
);

CREATE TABLE IF NOT EXISTS "group_chat" (
    "group_id" INTEGER NOT NULL,
    "from_id" INTEGER NOT NULL,
    "message" TEXT NOT NULL,
    "timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id"),
    FOREIGN KEY ("from_id") REFERENCES "users" ("user_id")
);
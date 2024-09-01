CREATE TABLE IF NOT EXISTS "events" (
    "event_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "group_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "title" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "date" TEXT NOT NULL,
    "timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
);

CREATE TABLE IF NOT EXISTS "event_users" (
    "event_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "going" INTEGER NOT NULL,
    FOREIGN KEY ("event_id") REFERENCES "events" ("event_id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
);
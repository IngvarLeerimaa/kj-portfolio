CREATE TABLE IF NOT EXISTS "groups" (
    "group_id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "title" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "admin_id" INTEGER NOT NULL,
    "timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("admin_id") REFERENCES "users" ("user_id")
);

CREATE TABLE IF NOT EXISTS "group_user" (
    "group_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "confirmed" INTEGER NOT NULL,
    FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id")
);

CREATE TABLE IF NOT EXISTS "group_post" (
    "post_id" INTEGER NOT NULL,
    "group_id" INTEGER NOT NULL,
    FOREIGN KEY ("post_id") REFERENCES "posts" ("post_id"),
    FOREIGN KEY ("group_id") REFERENCES "groups" ("group_id")
);
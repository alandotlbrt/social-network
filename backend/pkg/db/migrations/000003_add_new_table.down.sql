CREATE TABLE IF NOT EXISTS "chat"(
    "id_tchat" INTEGER PRIMARY KEY,
    "from_id" INT,
    "to_id" INT,
    "content" TEXT,
    "type" TEXT,
    "pictures" TEXT
);

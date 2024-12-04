CREATE TABLE IF NOT EXISTS "waiting_follow"(
    "id_who_waiting" INT,
    "id_who_accept"INT,
    PRIMARY KEY (id_who_waiting, id_who_accept),
    FOREIGN KEY (id_who_waiting) REFERENCES users(id_user),
    FOREIGN KEY (id_who_accept) REFERENCES users(id_user)
)

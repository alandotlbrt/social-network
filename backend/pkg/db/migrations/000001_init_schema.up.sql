CREATE TABLE IF NOT EXISTS "users" (
    "id_user" INTEGER PRIMARY KEY AUTOINCREMENT,
    "email" TEXT UNIQUE,
    "password" TEXT,
    "first_name" TEXT,
    "last_name" TEXT,
    "birthday" TEXT,
    "pp" TEXT,
    "nickname" TEXT UNIQUE,
    "about_me" TEXT,
    "privacy" TEXT,
    "follower" INT DEFAULT 0,
    "follow" INT DEFAULT 0,
    "cookie" TEXT
);


CREATE TABLE IF NOT EXISTS "follow" (
    "id_user" INT,
    "follow" INT,
    PRIMARY KEY (id_user, follow),
    FOREIGN KEY (id_user) REFERENCES users(id_user),
    FOREIGN KEY (follow) REFERENCES users(id_user)
);


CREATE TRIGGER IF NOT EXISTS increment_follow_count
AFTER INSERT ON follow
FOR EACH ROW
BEGIN
    UPDATE users
    SET follow = follow + 1
    WHERE id_user = NEW.follow;
END;

CREATE TRIGGER IF NOT EXISTS decrement_follow_count
AFTER DELETE ON follow
FOR EACH ROW
BEGIN
    UPDATE users
    SET follow = follow - 1
    WHERE id_user = OLD.follow;
END;


CREATE TRIGGER IF NOT EXISTS increment_followers_count
AFTER INSERT ON follow
FOR EACH ROW 
BEGIN 
    UPDATE users
    SET follower = follower + 1
    WHERE id_user = NEW.id_user;
END;


CREATE TRIGGER IF NOT EXISTS decrement_followers_count
AFTER DELETE ON follow
FOR EACH ROW 
BEGIN 
    UPDATE users
    SET follower = follower - 1
    WHERE id_user = OLD.id_user;
END;



 CREATE TABLE IF NOT EXISTS "post" (
    "id_post" INTEGER PRIMARY KEY,
    "id_user" INT,
    "content" TEXT,
    "image" TEXT,
    "privacy" TEXT,
    "like" int
);

CREATE TABLE IF NOT EXISTS "like"(
    "id_like" INTEGER PRIMARY KEY,
    "id_post" INT,
   "id_user" INT
);

 CREATE TABLE IF NOT EXISTS "allowed" (
    "id_post" INT,
    "id_user" INT
); 

CREATE TABLE IF NOT EXISTS "comment"(
    "id_comment" INTEGER PRIMARY KEY,
    "id_user" INT,
    "id_post" INT,
    "image" TEXT,
    "content" TEXT
);

CREATE TABLE IF NOT EXISTS "notifications"(
    "id_notification" INTEGER PRIMARY KEY,
    "user_id" INT,
    "content" TEXT,
    "type" TEXT
);

CREATE TABLE IF NOT EXISTS "chat"(
    "id_tchat" INTEGER PRIMARY KEY,
    "from_id" INT,
    "to_id" INT,
    "content" TEXT,
    "type" TEXT
);

CREATE TABLE IF NOT EXISTS "group"(
    "id_group" INTEGER PRIMARY KEY,
    "name" TEXT,
    "id_user" INT,
    "image" TEXT
);

CREATE TABLE IF NOT EXISTS "user_group"(
    "id_user_group" INTEGER PRIMARY KEY,
    "id_group" INT,
    "id_user" INT,
    "statue" TEXT
); 



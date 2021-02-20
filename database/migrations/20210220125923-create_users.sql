
-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
    user_id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL ,
    email VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS users;
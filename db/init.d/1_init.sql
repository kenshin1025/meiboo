CREATE DATABASE IF NOT EXISTS meiboo;

USE meiboo;

CREATE TABLE IF NOT EXISTS user (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  token VARCHAR(100) NOT NULL,
  created_at timestamp NOT NULL default current_timestamp,
  updated_at timestamp NOT NULL default current_timestamp on update current_timestamp
);

CREATE TABLE IF NOT EXISTS workspace (
  id int NOT NULL PRIMARY KEY,
  token VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS member (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  token VARCHAR(100) NOT NULL,
  image VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL,
  comment TEXT NOT NULL,
  workspace_id int NOT NULL,
  FOREIGN KEY (workspace_id) REFERENCES workspace(id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO
  workspace (id, token, name)
VALUES
  (
    1,
    "1111111-1111-1111-1111-111111111111",
    "システム工学研究会"
  );

INSERT INTO 
  member (token, image, name, comment, workspace_id)
VALUES
  ("1111111-1111-1111-1111-111111111111","jack", "健心", "こんにちは", 1),
  ("2222222-2222-2222-2222-222222222222","jeane", "けんしん", "コメントですよ", 1),
  ("3333333-3333-3333-3333-333333333333","jodi", "ああああああああああああ", "あああああああああああああああああああああああああああああ", 1);
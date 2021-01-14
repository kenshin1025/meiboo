CREATE DATABASE IF NOT EXISTS meiboo;

USE meiboo;

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

CREATE TABLE IF NOT EXISTS tag (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS member_tag (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  member_id int NOT NULL,
  tag_id int NOT NULL,
  FOREIGN KEY (member_id) REFERENCES member(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS workspace_tag (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  workspace_id int NOT NULL,
  tag_id int NOT NULL,
  FOREIGN KEY (workspace_id) REFERENCES workspace(id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE ON UPDATE CASCADE
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
  member (id, token, image, name, comment, workspace_id)
VALUES
  (
    1,
    "1111111-1111-1111-1111-111111111111",
    "jack",
    "健心",
    "こんにちは",
    1
  ),
  (
    2,
    "2222222-2222-2222-2222-222222222222",
    "jeane",
    "けんしん",
    "コメントですよ",
    1
  ),
  (
    3,
    "3333333-3333-3333-3333-333333333333",
    "jodi",
    "ああああああああああああ",
    "あああああああああああああああああああああああああああああ",
    1
  );

INSERT INTO
  tag (name)
VALUES
  ("1年生"),
  ("2年生"),
  ("3年生"),
  ("4年生"),
  ("プログラミング"),
  ("デザイン"),
  ("ゲーム"),
  ("Webアプリ"),
  ("JavaScript"),
  ("Ruby"),
  ("PHP"),
  ("Python"),
  ("Golang"),
  ("コンシス"),
  ("メディア");

INSERT INTO
  member_tag (member_id, tag_id)
VALUES
  (1, 1),
  (2, 2),
  (3, 3),
  (3, 5),
  (3, 7);
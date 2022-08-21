CREATE TABLE participants (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  uuid BIGINT UNIQUE,
  name VARCHAR,
  photo VARCHAR,
  description VARCHAR,
  votes INT DEFAULT(0)
);

CREATE TABLE voters (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  uuid BIGINT UNIQUE,
  participants BIGINT,
  FOREIGN KEY (participants) REFERENCES participants(uuid)
);
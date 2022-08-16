CREATE TABLE participants (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  uuid BIGINT UNIQUE,
  nickname VARCHAR,
  photo VARCHAR,
  information VARCHAR,
  votes INT DEFAULT(0)
);
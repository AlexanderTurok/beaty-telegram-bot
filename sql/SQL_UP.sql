CREATE TABLE participants (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  uuid BIGINT UNIQUE,
  name VARCHAR,
  photo VARCHAR,
  description VARCHAR,
  votes INT DEFAULT(0)
);
CREATE TABLE participants (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  uuid INT UNIQUE,
  nickaname VARCHAR,
  photo VARCHAR,
  information VARCHAR,
  votes INT DEFAULT(0)
);
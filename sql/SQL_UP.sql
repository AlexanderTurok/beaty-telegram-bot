CREATE TABLE participant (
  uuid VARCHAR(128) PRIMARY KEY,
  name TEXT,
  photo TEXT,
  description TEXT,
  likes INT DEFAULT(0)
);

CREATE TABLE voter (uuid VARCHAR(128) PRIMARY KEY);

CREATE TABLE voters_participant (
  voter_uuid VARCHAR(128) REFERENCES voter(uuid),
  participant_uuid VARCHAR(128) REFERENCES participant(uuid),
  PRIMARY KEY (voter_uuid, participant_uuid)
);
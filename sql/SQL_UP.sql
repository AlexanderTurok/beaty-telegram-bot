CREATE TABLE participant (
  uuid TEXT PRIMARY KEY,
  name TEXT,
  photo TEXT,
  description TEXT,
  likes INT DEFAULT(0)
);

CREATE TABLE voter (uuid TEXT PRIMARY KEY);

CREATE TABLE voters_participant (
  voter_uuid TEXT REFERENCES voter(uuid),
  participant_uuid TEXT REFERENCES participant(uuid),
  PRIMARY KEY (voter_uuid, participant_uuid)
);
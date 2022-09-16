CREATE TABLE participant (
  uuid VARCHAR(128) PRIMARY KEY,
  name TEXT,
  photo TEXT,
  description TEXT,
  likes INT DEFAULT(0)
);

CREATE TABLE voter (uuid VARCHAR(128) PRIMARY KEY);

CREATE TABLE voter_participant (
  voter_uuid VARCHAR(128) REFERENCES voter(uuid),
  participant_uuid VARCHAR(128) REFERENCES participant(uuid),
  PRIMARY KEY (voter_uuid, participant_uuid)
);

-- when create new participant
INSERT INTO
  voter_participant (participant_uuid, voter_uuid)
SELECT
  '1',
  uuid
FROM
  voter
WHERE
  voter.uuid <> '1';

-- when create new voter
INSERT INTO
  voter_participant (voter_uuid, participant_uuid)
SELECT
  '1',
  uuid
FROM
  participant
WHERE
  participant.uuid <> '1';

SELECT
  participant_uuid
FROM
  voter_participant
  LEFT JOIN participant ON participant.uuid = voter_participant.participant_uuid
WHERE
  voter_participant.voter_uuid = '1'
LIMIT
  1;
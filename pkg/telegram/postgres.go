package telegram

import "fmt"

func (b *Bot) isParticipantInDB(uuid int) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM participants WHERE uuid = %d);", uuid)
	err := b.postgres.QueryRow(query).Scan(&exists)

	return exists, err
}

func (b *Bot) getParticipantFromDB(uuid int) (*Participant, error) {
	var p Participant
	query := fmt.Sprintf("SELECT * FROM participants WHERE uuid=%d;", uuid)
	err := b.postgres.QueryRow(query).Scan(&p.Id, &p.Uuid, &p.Nickname, &p.Photo, &p.Information, &p.Votes)

	return &p, err
}

func (b *Bot) getAllParticipantsFromDB() (*[]Participant, error) {
	var pArray []Participant
	rows, err := b.postgres.Query("SELECT * FROM participants;")
	for rows.Next() {
		var p Participant
		if err := rows.Scan(&p.Id, &p.Uuid, &p.Nickname, &p.Photo, &p.Information, &p.Votes); err != nil {
			return &pArray, err
		}
		pArray = append(pArray, p)
	}
	return &pArray, err
}

func (b *Bot) addParticipantToDB(column string, row interface{}) error {
	query := fmt.Sprintf("INSERT INTO participants (%s) VALUES (%v);", column, row)
	_, err := b.postgres.Exec(query)

	return err
}

func (b *Bot) updateParticipantInDB(column, row string, uuid int) error {
	query := fmt.Sprintf("UPDATE participants SET %s='%s' WHERE uuid=%d", column, row, uuid)
	_, err := b.postgres.Exec(query)

	return err
}

func (b *Bot) updateVotesInDB(uuid string) error {
	query := fmt.Sprintf("UPDATE participants SET votes = votes + 1 WHERE uuid=%s", uuid)
	_, err := b.postgres.Exec(query)

	return err
}

func (b *Bot) deleteParticipantFromDB(column string, row interface{}) error {
	query := fmt.Sprintf("DELETE FROM participants WHERE %s=%v;", column, row)
	_, err := b.postgres.Exec(query)

	return err
}

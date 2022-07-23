package data

type Participant struct {
	Id          int
	Uuid        int
	Name        string
	Photo       string
	Description string
	Votes       int
}

func AddParticipantsToDatabase() (participants []Participant, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, photo, description, votes FROM particiapants")
	if err != nil {
		return
	}
	for rows.Next() {
		pt := Participant{}
		if err = rows.Scan(&pt.Id, &pt.Uuid, &pt.Name, &pt.Photo, &pt.Description, &pt.Votes); err != nil {
			return
		}
		participants = append(participants, pt)
	}
	rows.Close()
	return
}

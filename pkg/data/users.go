package data

type Participant struct {
	Id          int
	Uuid        int
	Name        string
	Photo       string
	Description string
	Votes       int
}

func IsParticipantRowExists(query string, args ...interface{}) (bool, error) {
	var exists bool
	err := Db.QueryRow(query, args...).Scan(&exists)
	return exists, err
}

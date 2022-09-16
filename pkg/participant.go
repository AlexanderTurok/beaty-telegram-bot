package telegram

import "errors"

type Participant struct {
	Uuid        string      `db:"participant_uuid"`
	Name        interface{} `db:"name"`
	Photo       interface{} `db:"photo"`
	Description interface{} `db:"description"`
	Likes       int         `db:"likes"`
}

func (p *Participant) Validate() error {
	if p.Name == nil && p.Photo == nil && p.Description == nil {
		return errors.New("profile is empty")
	}

	if p.Name == nil {
		return errors.New("name is empty")
	}

	if p.Description == nil {
		return errors.New("description is empty")
	}

	if p.Photo == nil {
		return errors.New("photo file is missing")
	}

	return nil
}

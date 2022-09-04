package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v9"
	_ "github.com/lib/pq"
)

type VoterRepository struct {
	context context.Context
	db      *sql.DB
	redis   *redis.Client
}

func NewVoterRepository(context context.Context, db *sql.DB, redis *redis.Client) *VoterRepository {
	return &VoterRepository{
		context: context,
		db:      db,
		redis:   redis,
	}
}

func (v *VoterRepository) Create(uuid int) error {
	query := fmt.Sprintf("INSERT INTO voters (uuid) VALUES (%d);", uuid)
	_, err := v.db.Exec(query)

	return err
}

// FIXME: fix
func (v *VoterRepository) GetParticipantUUID(uuid int) ([]int, error) {
	var p []int
	query := fmt.Sprintf("SELECT participants FROM voters where uuid='%d';", uuid)
	err := v.db.QueryRow(query).Scan(&p)

	return p, err
}

func (v *VoterRepository) DeleteParticipant(voterUuid, participantUuid int) error {
	query := fmt.Sprintf("")
	_, err := v.db.Exec(query)

	return err
}

func (v *VoterRepository) IsExists(uuid int) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM voters WHERE uuid='%d');", uuid)
	err := v.db.QueryRow(query).Scan(&exists)

	return exists, err
}

package repository

import (
	"context"
	"database/sql"

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

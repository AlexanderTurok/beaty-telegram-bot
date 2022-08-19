package repository

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type ParticipantRD struct {
	redis   *redis.Client
	context context.Context
}

func NewParticipantRD(context context.Context, redis *redis.Client) *ParticipantRD {
	return &ParticipantRD{
		redis:   redis,
		context: context,
	}
}

func (p *ParticipantRD) SetCache(uuid, value string) error {
	err := p.redis.Set(p.context, uuid, value, 0)
	return err.Err()
}

func (p *ParticipantRD) GetCache(uuid string) (string, error) {
	value, err := p.redis.Get(p.context, uuid).Result()
	return value, err
}

func (p *ParticipantRD) DeleteCache(uuid string) error {
	err := p.redis.Del(p.context, uuid)
	return err.Err()
}

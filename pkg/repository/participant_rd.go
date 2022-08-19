package repository

import (
	"context"
	"fmt"

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

func (p *ParticipantRD) SetCache(uuid int, value string) error {
	err := p.redis.Set(p.context, fmt.Sprint(uuid), value, 0)
	return err.Err()
}

func (p *ParticipantRD) GetCache(uuid int) (string, error) {
	value, err := p.redis.Get(p.context, fmt.Sprint(uuid)).Result()
	return value, err
}

func (p *ParticipantRD) DeleteCache(uuid int) error {
	err := p.redis.Del(p.context, fmt.Sprint(uuid))
	return err.Err()
}

package repository

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type VoterRD struct {
	redis   *redis.Client
	context context.Context
}

func NewVoterRD(context context.Context, redis *redis.Client) *VoterRD {
	return &VoterRD{
		redis: redis,
	}
}

func (v *VoterRD) SetCache(uuid string, value string) error {
	err := v.redis.Set(v.context, uuid, value, 0)
	return err.Err()
}

func (v *VoterRD) GetCache(uuid string) (string, error) {
	value, err := v.redis.Get(v.context, uuid).Result()
	return value, err
}

func (v *VoterRD) DeleteCache(uuid string) error {
	err := v.redis.Del(v.context, uuid)
	return err.Err()
}

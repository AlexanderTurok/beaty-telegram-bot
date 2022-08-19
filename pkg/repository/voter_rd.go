package repository

import (
	"context"
	"fmt"

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

func (v *VoterRD) SetCache(uuid int, value string) error {
	err := v.redis.Set(v.context, fmt.Sprint(uuid), value, 0)
	return err.Err()
}

func (v *VoterRD) GetCache(uuid int) (string, error) {
	value, err := v.redis.Get(v.context, fmt.Sprint(uuid)).Result()
	return value, err
}

func (v *VoterRD) DeleteCache(uuid int) error {
	err := v.redis.Del(v.context, fmt.Sprint(uuid))
	return err.Err()
}

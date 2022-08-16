package telegram

import "fmt"

func (b *Bot) setCache(key int, value interface{}) error {
	err := b.redis.Set(b.ctx, fmt.Sprint(key), value, 0)
	return err.Err()
}

func (b *Bot) getCache(key int) string {
	value, _ := b.redis.Get(b.ctx, fmt.Sprint(key)).Result()
	return value
}

func (b *Bot) deleteCache(key int) error {
	err := b.redis.Del(b.ctx, fmt.Sprint(key))
	return err.Err()
}

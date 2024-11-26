package infrastructurepinger

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisPinger struct {
	rds *redis.Client
}

func NewRedisPinger(rds *redis.Client) *RedisPinger {
	return &RedisPinger{rds: rds}
}

func (p *RedisPinger) Ping(ctx context.Context) error {
	_, err := p.rds.Ping(ctx).Result()
	return err
}

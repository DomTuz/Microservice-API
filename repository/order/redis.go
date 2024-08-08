package order

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/DomTuz/Microservice-API/model"
)

type RedisRepo struct {
	Client *redis.Client
}

func (r *RedisRepo) Insert(ctx context.Context, order model.Order) error {
	return nil
}
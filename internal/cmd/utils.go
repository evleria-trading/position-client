package cmd

import (
	"context"
	"fmt"
	"github.com/evleria/position-client/internal/config"
	"github.com/evleria/position-service/protocol/pb"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func getRedis(cfg *config.Сonfig) *redis.Client {
	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPass,
	}

	redisClient := redis.NewClient(opts)
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	return redisClient
}

func getClient(cfg *config.Сonfig) pb.PositionServiceClient {
	conn, err := grpc.Dial(cfg.PositionServiceUrl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	return pb.NewPositionServiceClient(conn)
}

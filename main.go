package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evleria/position-client/internal/cache"
	"github.com/evleria/position-client/internal/cmd"
	"github.com/evleria/position-client/internal/config"
	"github.com/evleria/position-client/internal/consumer"
	"github.com/evleria/position-service/protocol/pb"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"strings"
)

func main() {
	cfg := getConfig()
	redisClient := getRedis(cfg)
	grpcClient := getClient(cfg)

	pricesCache := cache.NewPriceCache()
	priceConsumer := consumer.NewPriceConsumer(redisClient)
	prices := priceConsumer.Consume(context.Background())
	go func() {
		for price := range prices {
			pricesCache.UpdatePrice(*price)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Split(input, " ")
		rootCmd := cmd.NewRootCmd(grpcClient, pricesCache)
		rootCmd.SetArgs(args)
		err := rootCmd.Execute()
		if err != nil {
			log.Error(err)
		}
	}
}

func getConfig() *config.Сonfig {
	cfg := new(config.Сonfig)
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

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

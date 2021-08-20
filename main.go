package main

import (
	"bufio"
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/evleria/position-client/internal/cache"
	"github.com/evleria/position-client/internal/cmd"
	"github.com/evleria/position-client/internal/config"
	"github.com/evleria/position-client/internal/consumer"
	positionPb "github.com/evleria/position-service/protocol/pb"
	pricePb "github.com/evleria/price-service/protocol/pb"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"strings"
)

func main() {
	cfg := getConfig()
	positionGrpcClient := getPositionGrpcClient(cfg)
	getPriceGrpcClient := getPriceGrpcClient(cfg)

	pricesCache := cache.NewPriceCache()
	priceConsumer := consumer.NewPriceConsumer(getPriceGrpcClient)
	prices, err := priceConsumer.Consume(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for price := range prices {
			err := pricesCache.UpdatePrice(price)
			if err != nil {
				log.Error(err)
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		args := strings.Split(input, " ")
		rootCmd := cmd.NewRootCmd(positionGrpcClient, pricesCache)
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

func getPositionGrpcClient(cfg *config.Сonfig) positionPb.PositionServiceClient {
	conn, err := grpc.Dial(cfg.PositionServiceUrl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	return positionPb.NewPositionServiceClient(conn)
}

func getPriceGrpcClient(cfg *config.Сonfig) pricePb.PriceServiceClient {
	conn, err := grpc.Dial(cfg.PriceServiceUrl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	return pricePb.NewPriceServiceClient(conn)
}

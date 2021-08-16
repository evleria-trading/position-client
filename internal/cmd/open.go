package cmd

import (
	"context"
	"github.com/evleria/position-client/internal/config"
	"github.com/evleria/position-client/internal/consumer"
	"github.com/evleria/position-client/internal/model"
	"github.com/evleria/position-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type OpenPositionCmdOptions struct {
	Config    *config.Сonfig
	Symbol    string
	IsBuyType bool
}

func NewOpenPositionCmd(cfg *config.Сonfig) *cobra.Command {
	opts := &OpenPositionCmdOptions{
		Config: cfg,
	}

	openCmd := &cobra.Command{
		Use: "open",
		RunE: func(c *cobra.Command, args []string) error {
			return runOpen(opts)
		},
	}

	openCmd.Flags().StringVarP(&opts.Symbol, "symbol", "s", "", "defines symbol for desired position")
	openCmd.Flags().BoolVarP(&opts.IsBuyType, "buy", "b", false, "defines type for desired position")
	_ = openCmd.MarkFlagRequired("symbol")

	return openCmd
}

func runOpen(opts *OpenPositionCmdOptions) error {
	redisClient := getRedis(opts.Config)
	priceConsumer := consumer.NewPriceConsumer(redisClient)

	grpcClient := getClient(opts.Config)

	ctx, cancel := context.WithCancel(context.Background())
	prices := priceConsumer.Consume(ctx)

	var price *model.Price
	for p := range prices {
		if p.Symbol == opts.Symbol {
			price = p
			cancel()
		}
	}

	request := &pb.OpenPositionRequest{
		Symbol:    opts.Symbol,
		IsBuyType: opts.IsBuyType,
		PriceId:   price.Id,
	}
	response, err := grpcClient.OpenPosition(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{"id": response.PositionId}).Info("Opened position")
	return nil
}

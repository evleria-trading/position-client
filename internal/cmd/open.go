package cmd

import (
	"context"
	"github.com/evleria/position-client/internal/cache"
	positionPb "github.com/evleria/position-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type OpenPositionCmdOptions struct {
	Symbol    string
	IsBuyType bool
}

func NewOpenPositionCmd(grpcClient positionPb.PositionServiceClient, pricesCache cache.Price) *cobra.Command {
	opts := new(OpenPositionCmdOptions)

	openCmd := &cobra.Command{
		Use:  "open [OPTIONS] SYMBOL",
		Args: cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			opts.Symbol = args[0]
			return runOpen(opts, grpcClient, pricesCache)
		},
	}

	openCmd.Flags().BoolVarP(&opts.IsBuyType, "buy", "b", false, "defines type for desired position")

	return openCmd
}

func runOpen(opts *OpenPositionCmdOptions, grpcClient positionPb.PositionServiceClient, pricesCache cache.Price) error {
	price, err := pricesCache.GetPrice(opts.Symbol)
	if err != nil {
		return err
	}

	request := &positionPb.OpenPositionRequest{
		Symbol:    opts.Symbol,
		IsBuyType: opts.IsBuyType,
		PriceId:   price.Id,
	}
	response, err := grpcClient.OpenPosition(context.Background(), request)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"id": response.PositionId}).Info("Opened position")
	return nil
}

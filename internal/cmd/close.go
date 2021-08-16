package cmd

import (
	"context"
	"github.com/evleria/position-client/internal/cache"
	"github.com/evleria/position-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

type ClosePositionCmdOptions struct {
	PositionId int64
}

func NewClosePositionCmd(grpcClient pb.PositionServiceClient, pricesCache cache.Price) *cobra.Command {
	opts := new(ClosePositionCmdOptions)

	closeCmd := &cobra.Command{
		Use:  "close POSITION_ID",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			opts.PositionId = int64(id)
			return runClose(opts, grpcClient, pricesCache)
		},
	}

	return closeCmd
}

func runClose(opts *ClosePositionCmdOptions, grpcClient pb.PositionServiceClient, pricesCache cache.Price) error {
	position, err := grpcClient.GetOpenPosition(context.Background(), &pb.GetOpenPositionRequest{
		PositionId: opts.PositionId,
	})
	if err != nil {
		return err
	}

	price, err := pricesCache.GetPrice(position.Symbol)
	if err != nil {
		return err
	}

	response, err := grpcClient.ClosePosition(context.Background(), &pb.ClosePositionRequest{
		PositionId: opts.PositionId,
		PriceId:    price.Id,
	})
	log.WithFields(log.Fields{"id": opts.PositionId, "profit": response.Profit}).Info("Closed position")
	return nil
}

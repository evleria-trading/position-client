package cmd

import (
	"context"
	"github.com/evleria/position-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

type TakeProfitCmdOptions struct {
	PositionId int64
	TakeProfit float64
}

func NewSetTakeProfitCmd(grpcClient pb.PositionServiceClient) *cobra.Command {
	opts := new(TakeProfitCmdOptions)

	takeCmd := &cobra.Command{
		Use:  "tp POSITION_ID VALUE",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			v, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				return err
			}
			opts.PositionId, opts.TakeProfit = int64(id), v
			return runSetTakeProfit(opts, grpcClient)
		},
	}
	return takeCmd
}

func runSetTakeProfit(opts *TakeProfitCmdOptions, grpcClient pb.PositionServiceClient) error {
	_, err := grpcClient.SetTakeProfit(context.Background(), &pb.SetTakeProfitRequest{
		PositionId: opts.PositionId,
		TakeProfit: opts.TakeProfit,
	})
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"id": opts.PositionId, "take_profit": opts.TakeProfit}).Info("Set take profit")
	return nil
}

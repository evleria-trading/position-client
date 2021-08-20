package cmd

import (
	"context"
	positionPb "github.com/evleria/position-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

type SetStopLossCmdOptions struct {
	PositionId int64
	StopLoss   float64
}

func NewSetStopLossCmd(grpcClient positionPb.PositionServiceClient) *cobra.Command {
	opts := new(SetStopLossCmdOptions)

	setCmd := &cobra.Command{
		Use:  "sl POSITION_ID VALUE",
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
			opts.PositionId, opts.StopLoss = int64(id), -v
			return runSetStopLoss(opts, grpcClient)
		},
	}

	return setCmd
}

func runSetStopLoss(opts *SetStopLossCmdOptions, grpcClient positionPb.PositionServiceClient) error {
	_, err := grpcClient.SetStopLoss(context.Background(), &positionPb.SetStopLossRequest{
		PositionId: opts.PositionId,
		StopLoss:   opts.StopLoss,
	})
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"id": opts.PositionId, "stop_loss": opts.StopLoss}).Info("Set stop loss")
	return nil
}

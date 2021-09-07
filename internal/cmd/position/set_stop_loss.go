package position

import (
	"context"
	"github.com/evleria-trading/position-client/internal/scope"
	positionPb "github.com/evleria-trading/position-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

type SetStopLossCmdOptions struct {
	PositionId int64
	StopLoss   float64
}

func NewSetStopLossCmd(s *scope.Scope) *cobra.Command {
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
			return runSetStopLoss(opts, s)
		},
	}

	return setCmd
}

func runSetStopLoss(opts *SetStopLossCmdOptions, s *scope.Scope) error {
	_, err := s.PositionsClient.SetStopLoss(context.Background(), &positionPb.SetStopLossRequest{
		PositionId: opts.PositionId,
		StopLoss:   opts.StopLoss,
		UserId:     s.CurrentUserId,
	})
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"id": opts.PositionId, "stop_loss": opts.StopLoss}).Info("Set stop loss")
	return nil
}

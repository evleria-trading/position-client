package position

import (
	"context"
	"github.com/evleria-trading/position-client/internal/scope"
	positionPb "github.com/evleria-trading/position-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

type TakeProfitCmdOptions struct {
	PositionId int64
	TakeProfit float64
}

func NewSetTakeProfitCmd(s *scope.Scope) *cobra.Command {
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
			return runSetTakeProfit(opts, s)
		},
	}
	return takeCmd
}

func runSetTakeProfit(opts *TakeProfitCmdOptions, s *scope.Scope) error {
	_, err := s.PositionsClient.SetTakeProfit(context.Background(), &positionPb.SetTakeProfitRequest{
		PositionId: opts.PositionId,
		TakeProfit: opts.TakeProfit,
	})
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"id": opts.PositionId, "take_profit": opts.TakeProfit}).Info("Set take profit")
	return nil
}

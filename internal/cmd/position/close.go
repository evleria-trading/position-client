package position

import (
	"context"
	"github.com/evleria-trading/position-client/internal/scope"
	positionPb "github.com/evleria-trading/position-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

type ClosePositionCmdOptions struct {
	PositionId int64
}

func NewClosePositionCmd(s *scope.Scope) *cobra.Command {
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
			return runClose(opts, s)
		},
	}

	return closeCmd
}

func runClose(opts *ClosePositionCmdOptions, s *scope.Scope) error {
	position, err := s.PositionsClient.GetOpenPosition(context.Background(), &positionPb.GetOpenPositionRequest{
		PositionId: opts.PositionId,
	})
	if err != nil {
		return err
	}

	price, err := s.PricesCache.GetPrice(position.Symbol)
	if err != nil {
		return err
	}

	response, err := s.PositionsClient.ClosePosition(context.Background(), &positionPb.ClosePositionRequest{
		PositionId: opts.PositionId,
		PriceId:    price.Id,
	})
	log.WithFields(log.Fields{"id": opts.PositionId, "profit": response.Profit}).Info("Closed position")
	return nil
}

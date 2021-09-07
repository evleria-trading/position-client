package position

import (
	"context"
	"github.com/evleria-trading/position-client/internal/scope"
	positionPb "github.com/evleria-trading/position-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type OpenPositionCmdOptions struct {
	Symbol    string
	IsBuyType bool
}

func NewOpenPositionCmd(s *scope.Scope) *cobra.Command {
	opts := new(OpenPositionCmdOptions)

	openCmd := &cobra.Command{
		Use:  "open [OPTIONS] SYMBOL",
		Args: cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			opts.Symbol = args[0]
			return runOpen(opts, s)
		},
	}

	openCmd.Flags().BoolVarP(&opts.IsBuyType, "buy", "b", false, "defines type for desired position")

	return openCmd
}

func runOpen(opts *OpenPositionCmdOptions, s *scope.Scope) error {
	price, err := s.PricesCache.GetPrice(opts.Symbol)
	if err != nil {
		return err
	}

	request := &positionPb.OpenPositionRequest{
		Symbol:    opts.Symbol,
		IsBuyType: opts.IsBuyType,
		PriceId:   price.Id,
		UserId:    1,
	}
	response, err := s.PositionsClient.OpenPosition(context.Background(), request)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"id": response.PositionId}).Info("Opened position")
	return nil
}

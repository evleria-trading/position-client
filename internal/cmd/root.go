package cmd

import (
	"github.com/evleria/position-client/internal/cache"
	positionPb "github.com/evleria/position-service/protocol/pb"
	"github.com/spf13/cobra"
)

func NewRootCmd(grpcClient positionPb.PositionServiceClient, pricesCache cache.Price) *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(NewOpenPositionCmd(grpcClient, pricesCache))
	rootCmd.AddCommand(NewClosePositionCmd(grpcClient, pricesCache))
	rootCmd.AddCommand(NewSetStopLossCmd(grpcClient))
	rootCmd.AddCommand(NewSetTakeProfitCmd(grpcClient))

	return rootCmd
}

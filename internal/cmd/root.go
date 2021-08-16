package cmd

import (
	"github.com/evleria/position-client/internal/config"
	"github.com/spf13/cobra"
)

func NewRootCmd(cfg *config.Сonfig) *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(NewOpenPositionCmd(cfg))
	rootCmd.AddCommand(NewClosePositionCmd(cfg))

	return rootCmd
}

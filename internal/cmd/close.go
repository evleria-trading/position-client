package cmd

import (
	"github.com/evleria/position-client/internal/config"
	"github.com/spf13/cobra"
)

func NewClosePositionCmd(*config.Сonfig) *cobra.Command {
	closeCmd := &cobra.Command{
		Use: "close",
	}

	return closeCmd
}

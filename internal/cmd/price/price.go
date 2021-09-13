package price

import (
	"github.com/evleria-trading/position-client/internal/scope"
	"github.com/spf13/cobra"
)

func NewPriceCmd(s *scope.Scope) *cobra.Command {
	cmd := &cobra.Command{
		Use: "price",
	}

	cmd.AddCommand(NewListCmd(s))

	return cmd
}

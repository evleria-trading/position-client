package balance

import (
	"github.com/evleria-trading/position-client/internal/scope"
	"github.com/spf13/cobra"
)

func NewBalanceCmd(s *scope.Scope) *cobra.Command {
	cmd := &cobra.Command{
		Use: "balance",
	}
	cmd.AddCommand(NewGetBalanceCmd(s))
	cmd.AddCommand(NewSetBalanceCmd(s))

	return cmd
}

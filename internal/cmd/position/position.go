package position

import (
	"github.com/evleria-trading/position-client/internal/scope"
	"github.com/spf13/cobra"
)

func NewPositionCmd(s *scope.Scope) *cobra.Command {
	cmd := &cobra.Command{
		Use: "pos",
	}

	cmd.AddCommand(NewOpenPositionCmd(s))
	cmd.AddCommand(NewClosePositionCmd(s))
	cmd.AddCommand(NewSetStopLossCmd(s))
	cmd.AddCommand(NewSetTakeProfitCmd(s))

	return cmd
}

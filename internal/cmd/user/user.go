package user

import (
	"github.com/evleria-trading/position-client/internal/scope"
	"github.com/spf13/cobra"
)

func NewUserCmd(s *scope.Scope) *cobra.Command {
	cmd := &cobra.Command{
		Use: "user",
	}
	cmd.AddCommand(NewSetUserCmd(s))
	cmd.AddCommand(NewCreateCmd(s))

	return cmd
}

package cmd

import (
	"github.com/evleria-trading/position-client/internal/cmd/position"
	"github.com/evleria-trading/position-client/internal/cmd/user"
	"github.com/evleria-trading/position-client/internal/scope"
	"github.com/spf13/cobra"
)

func NewRootCmd(s *scope.Scope) *cobra.Command {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(user.NewUserCmd(s))
	if s.IsUserSet() {
		rootCmd.AddCommand(position.NewPositionCmd(s))
	}

	return rootCmd
}

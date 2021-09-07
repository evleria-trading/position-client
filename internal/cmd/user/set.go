package user

import (
	"fmt"
	"github.com/evleria-trading/position-client/internal/scope"
	"github.com/spf13/cobra"
	"strconv"
)

type SetUserCmdOptions struct {
	UserId int64
}

func NewSetUserCmd(s *scope.Scope) *cobra.Command {
	setCmd := &cobra.Command{
		Use:  "set USER_ID",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			opts := &SetUserCmdOptions{UserId: int64(id)}
			return runSetUser(s, opts)
		},
	}
	return setCmd
}

func runSetUser(s *scope.Scope, opts *SetUserCmdOptions) error {
	err := s.SetUserId(opts.UserId)
	if err != nil {
		return err
	}
	fmt.Println("userID successfully set to", opts.UserId)
	return nil
}

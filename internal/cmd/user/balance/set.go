package balance

import (
	"context"
	"github.com/evleria-trading/position-client/internal/scope"
	userPb "github.com/evleria-trading/user-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

type SetBalanceCmdOptions struct {
	Balance float64
}

func NewSetBalanceCmd(s *scope.Scope) *cobra.Command {
	setCmd := &cobra.Command{
		Use:  "set",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			v, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				return err
			}
			opts := &SetBalanceCmdOptions{
				Balance: v,
			}
			return runSetBalance(s, opts)

		},
	}
	return setCmd
}

func runSetBalance(s *scope.Scope, opts *SetBalanceCmdOptions) error {
	_, err := s.UserClient.SetBalance(context.Background(), &userPb.SetBalanceRequest{
		UserId:  s.CurrentUserId,
		Balance: opts.Balance})
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"balance": opts.Balance}).Info("Set balance")
	return nil
}

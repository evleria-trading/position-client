package balance

import (
	"context"
	"fmt"
	"github.com/evleria-trading/position-client/internal/scope"
	userPb "github.com/evleria-trading/user-service/protocol/pb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewGetBalanceCmd(s *scope.Scope) *cobra.Command {
	getCmd := &cobra.Command{
		Use: "get",
		RunE: func(c *cobra.Command, _ []string) error {
			return runGetBalance(s)
		},
	}
	return getCmd
}

func runGetBalance(s *scope.Scope) error {
	request := &userPb.GetBalanceRequest{
		UserId: s.CurrentUserId,
	}
	response, err := s.UserClient.GetBalance(context.Background(), request)
	if err != nil {
		return err
	}
	fmt.Println("get started")
	log.WithFields(log.Fields{"balance": response.Balance}).Info("Got balance")
	return nil
}

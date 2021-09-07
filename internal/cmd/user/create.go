package user

import (
	"context"
	"github.com/evleria-trading/position-client/internal/scope"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCreateCmd(s *scope.Scope) *cobra.Command {
	createCmd := &cobra.Command{
		Use: "create",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCreate(s)
		},
	}
	return createCmd
}

func runCreate(s *scope.Scope) error {
	response, err := s.UserClient.CreateUser(context.Background(), &empty.Empty{})
	if err != nil {
		return err
	}

	err = s.SetUserId(response.UserId)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{"id": response.UserId}).Info("Created new user and set as current")
	return nil
}

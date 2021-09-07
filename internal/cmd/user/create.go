package user

import (
	"context"
	"fmt"
	"github.com/evleria-trading/position-client/internal/scope"
	"github.com/golang/protobuf/ptypes/empty"
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
	fmt.Println("userID successfully set to", response.UserId)
	return nil
}

package scope

import (
	"errors"
	"github.com/evleria-trading/position-client/internal/cache"
	positionPb "github.com/evleria-trading/position-service/protocol/pb"
	userPb "github.com/evleria-trading/user-service/protocol/pb"
)

type Scope struct {
	PositionsClient positionPb.PositionServiceClient
	UserClient      userPb.UserServiceClient
	PricesCache     cache.Price
	CurrentUserId   int64
}

func NewScope(positionsClient positionPb.PositionServiceClient, userClient userPb.UserServiceClient, pricesCache cache.Price) *Scope {
	return &Scope{
		PositionsClient: positionsClient,
		UserClient:      userClient,
		PricesCache:     pricesCache,
		CurrentUserId:   0,
	}
}

func (s *Scope) IsUserSet() bool {
	return s.CurrentUserId > 0
}

func (s *Scope) SetUserId(id int64) error {
	if id <= 0 {
		return errors.New("id is not valid")
	}
	s.CurrentUserId = id
	return nil
}

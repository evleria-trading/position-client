package scope

import (
	"github.com/evleria-trading/position-client/internal/cache"
	positionPb "github.com/evleria-trading/position-service/protocol/pb"
)

type Scope struct {
	PositionsClient positionPb.PositionServiceClient
	PricesCache     cache.Price
}

func NewScope(positionsClient positionPb.PositionServiceClient, pricesCache cache.Price) *Scope {
	return &Scope{
		PositionsClient: positionsClient,
		PricesCache:     pricesCache,
	}
}

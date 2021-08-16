package cache

import (
	"errors"
	"github.com/evleria/position-client/internal/model"
	"sync"
)

var (
	ErrPriceNotFound = errors.New("price not found")
	ErrPriceIsNil    = errors.New("price is nil")
)

type Price interface {
	GetPrice(symbol string) (*model.Price, error)
	UpdatePrice(price *model.Price) error
}

type price struct {
	m  map[string]model.Price
	mx sync.RWMutex
}

func NewPriceCache() Price {
	return &price{
		m:  map[string]model.Price{},
		mx: sync.RWMutex{},
	}
}

func (p *price) GetPrice(symbol string) (*model.Price, error) {
	p.mx.RLock()
	defer p.mx.RUnlock()

	if v, ok := p.m[symbol]; ok {
		return &model.Price{
			Id:  v.Id,
			Ask: v.Ask,
			Bid: v.Bid,
		}, nil
	}
	return nil, ErrPriceNotFound
}

func (p *price) UpdatePrice(price *model.Price) error {
	p.mx.Lock()
	defer p.mx.Unlock()

	if price == nil {
		return ErrPriceIsNil
	}
	p.m[price.Symbol] = *price
	return nil
}

package cache

import (
	"errors"
	"github.com/evleria-trading/position-client/internal/model"
	"sort"
	"sync"
)

var (
	ErrPriceNotFound = errors.New("price not found")
	ErrPriceIsNil    = errors.New("price is nil")
)

type Price interface {
	GetPrice(symbol string) (*model.Price, error)
	UpdatePrice(price *model.Price) error
	GetPrices() []model.Price
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

func (p *price) GetPrices() []model.Price {
	p.mx.RLock()
	result := make([]model.Price, 0, len(p.m))
	for _, pr := range p.m {
		result = append(result, pr)
	}
	p.mx.RUnlock()

	sort.Slice(result, func(i, j int) bool {
		return result[i].Symbol < result[j].Symbol
	})
	return result
}

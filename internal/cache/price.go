package cache

import (
	"errors"
	"github.com/evleria/position-client/internal/model"
	"sync"
)

var ErrPriceNotFound = errors.New("price not found")

type Price interface {
	GetPrice(symbol string) (model.Price, error)
	UpdatePrice(price model.Price)
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

func (p *price) GetPrice(symbol string) (model.Price, error) {
	p.mx.RLock()
	defer p.mx.RUnlock()

	if v, ok := p.m[symbol]; ok {
		return v, nil
	}
	return model.Price{}, ErrPriceNotFound
}

func (p *price) UpdatePrice(price model.Price) {
	p.mx.Lock()
	defer p.mx.Unlock()

	p.m[price.Symbol] = price
}

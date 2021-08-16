package consumer

import (
	"context"
	"errors"
	"fmt"
	"github.com/evleria/position-client/internal/model"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Price interface {
	Consume(ctx context.Context) <-chan *model.Price
}

type price struct {
	redis *redis.Client
}

func NewPriceConsumer(redisClient *redis.Client) Price {
	return &price{
		redis: redisClient,
	}
}

func (p *price) Consume(ctx context.Context) <-chan *model.Price {
	id := fmt.Sprintf("%d000-0", time.Now().Unix())
	ch := make(chan *model.Price)

	go func() {
		for {
			args := &redis.XReadArgs{
				Streams: []string{"prices", id},
			}
			r, err := p.redis.XRead(ctx, args).Result()
			if err != nil {
				log.Error(err)
				continue
			}

			for _, message := range r[0].Messages {
				pr, err := decodeMessage(message)
				if err != nil {
					log.Error(err)
					break
				}

				log.WithFields(log.Fields{
					"id":     pr.Id,
					"symbol": pr.Symbol,
					"ask":    pr.Ask,
					"bid":    pr.Bid,
				}).Debug("Consumed price message")
				ch <- &pr

				id = message.ID
			}
		}
	}()

	return ch
}

func decodeMessage(message redis.XMessage) (model.Price, error) {
	symbol, err := decodeString(message.Values["symbol"])
	if err != nil {
		return model.Price{}, err
	}

	ask, err := decodeFloat64(message.Values["ask"])
	if err != nil {
		return model.Price{}, err
	}

	bid, err := decodeFloat64(message.Values["bid"])
	if err != nil {
		return model.Price{}, err
	}
	return model.Price{
		Id:     message.ID,
		Symbol: symbol,
		Ask:    ask,
		Bid:    bid,
	}, nil
}

func decodeString(v interface{}) (string, error) {
	if v == nil {
		return "", errors.New("cannot decode nil")
	}
	if str, ok := v.(string); ok {
		return str, nil
	}
	return "", errors.New("cannot convert to string")
}

func decodeFloat64(v interface{}) (float64, error) {
	str, err := decodeString(v)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(str, 64)
}

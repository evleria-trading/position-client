package consumer

import (
	"context"
	"github.com/evleria/position-client/internal/model"
	pricePb "github.com/evleria/price-service/protocol/pb"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
)

type Price interface {
	Consume(ctx context.Context) (<-chan *model.Price, error)
}

type price struct {
	priceClient pricePb.PriceServiceClient
}

func NewPriceConsumer(priceClient pricePb.PriceServiceClient) Price {
	return &price{
		priceClient: priceClient,
	}
}

func (p *price) Consume(ctx context.Context) (<-chan *model.Price, error) {
	stream, err := p.priceClient.GetPrices(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	ch := make(chan *model.Price)

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Error(err)
				return
			}

			pr := model.Price{
				Id:     msg.Id,
				Symbol: msg.Symbol,
				Ask:    msg.Ask,
				Bid:    msg.Bid,
			}

			log.WithFields(log.Fields{
				"id":     pr.Id,
				"symbol": pr.Symbol,
				"ask":    pr.Ask,
				"bid":    pr.Bid,
			}).Debug("Consumed price message")
			ch <- &pr
		}
	}()

	return ch, nil
}

package service

import (
	"Wb-L0/internal/model"
	"github.com/nats-io/stan.go"
	"log"
)

func (s *Service) StanSubscribe(channel string, cb func(msg *stan.Msg), opts ...stan.SubscriptionOption) (stan.Subscription, error) {
	sub, err := s.conn.Subscribe(channel, cb, opts...)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func (s *Service) StanHandler(msg *stan.Msg) {
	log.Println("Пришли данные...")
	order, err := model.JsonToOrder(msg.Data)
	if err != nil {
		log.Println(err)
		return
	}
	if err = s.store.Add(order); err != nil {
		log.Println(err)
		return
	}
	log.Println("UUID:", order.UUID)
}

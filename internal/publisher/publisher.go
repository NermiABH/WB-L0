package publisher

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	conn *stan.Conn
}

const (
	clusterID  = "test-cluster"
	clientID   = "test-publisher"
	chanel     = "foo"
	pathErrors = "testError"
)

func Start() error {
	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	conn, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	go func() {
		for {
			var s string
			fmt.Scan(&s)
			var order []byte
			if s != "auto" {
				continue
			}
			order, err = JsonFromRandom()
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(string(order))
			if err = conn.Publish(chanel, order); err != nil {
				log.Println(err)
			}
		}
	}()

	<-done
	return nil
}

func JsonFromFile(path string) ([]byte, error) {
	jsonOrder, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return jsonOrder, err
}

func JsonFromRandom() ([]byte, error) {
	order := NewFakeOrder()
	if err := gofakeit.Struct(&order); err != nil {
		return nil, err
	}
	jsonOrder, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	return jsonOrder, nil
}

package service

import (
	"Wb-L0/internal/store"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	pathErrors = "testError"
)

type Service struct {
	conn  stan.Conn
	store *store.Store
	route *http.ServeMux
}

func Start(config *Config) error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Connecting to nats server ...")
	conn, err := stan.Connect(config.Stan.ClusterID, config.Stan.ClientID)
	if err != nil {
		return err
	}
	defer conn.Close()
	str, err := store.NewStore(config.DatabaseURL())
	if err != nil {
		return err
	}
	defer str.Close()
	mux := http.NewServeMux()
	src := &Service{conn: conn, store: str, route: mux}
	src.route.Handle("/", http.FileServer(http.Dir("internal/service/static")))
	src.route.HandleFunc("/orders", src.GetByUUID)

	sb, err := src.StanSubscribe(config.Stan.Chanel, src.StanHandler, stan.StartWithLastReceived())
	if err != nil {
		return err
	}
	defer sb.Close()

	go func() {
		log.Printf("Starting server %s ...", config.ServerAddr())
		if err := http.ListenAndServe(config.ServerAddr(), src.route); err != nil {
			panic(err)
		}
	}()

	<-done
	return nil
}

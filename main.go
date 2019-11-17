package main

import (
	"fmt"
	"github.com/terawork-com/message-socket-service/pkg/config"
	"github.com/terawork-com/message-socket-service/pkg/domain"
	log "github.com/terawork-com/message-socket-service/pkg/log"
	"github.com/terawork-com/message-socket-service/pkg/handlers"
	"go.uber.org/zap"
	"net/http"
)

var pool *domain.Pool

func main() {
	// Load config
	conf, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	pool = domain.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
		return
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// define our WebSocket endpoint
func serveWs(pool *domain.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := handlers.Upgrade(w, r)
	if err != nil {
		log.Error(err)
	}
	client := &domain.Client{
		ID:   "",
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	defer func() { pool.Leave <- client }()
	go client.Write()
	client.Read()
	log.Info("New client connected!")
}
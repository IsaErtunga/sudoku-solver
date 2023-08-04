package src

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// Parameterize if needed
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type server struct {
	log     *log.Logger
	addr    string
	handler *http.ServeMux
}

func NewServer(serviceName, bindAdress string, log *log.Logger) *server {
	mux := http.NewServeMux()
	return &server{
		log:     log,
		addr:    bindAdress,
		handler: mux,
	}
}

func (s *server) InitServer(socketHandler *socketHandler) {
	s.handler.Handle("/live", socketHandler)

	httpServer := http.Server{
		Addr:         s.addr,
		Handler:      s.handler,
		ErrorLog:     s.log,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Printf("Starting server on port %s\n", s.addr)
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	serverChan := make(chan os.Signal, 1)
	signal.Notify(serverChan, os.Interrupt)
	signal.Notify(serverChan, os.Kill)

	// Block until a signal is received.
	sig := <-serverChan
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	httpServer.Shutdown(ctx)
}

package src

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type socketHandler struct {
	log *log.Logger
}

func NewSocketHandler(log *log.Logger) *socketHandler {
	return &socketHandler{log: log}
}

func reader() {}

func writer(conn *websocket.Conn) {
	for {
		err := conn.WriteMessage(1, []byte("hello"))
		if err != nil {
			panic("Panic!")
		}
	}
}

func (sh *socketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicln("Websocket handshake failed")
	}

	// Initialize & solve puzzle
	board := InitBoard()
	game := NewGame(board)
	game.Solve(BruteForce)

	go writer(conn)
}

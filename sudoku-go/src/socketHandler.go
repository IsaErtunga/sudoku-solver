package src

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type socketHandler struct {
	log *log.Logger
}

func NewSocketHandler(log *log.Logger) *socketHandler {
	return &socketHandler{log: log}
}

func reader(conn *websocket.Conn) {
	for {
		msgType, buf, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			panic("Error when reading")
		}
		fmt.Println(msgType, buf)
	}
}

func writer(conn *websocket.Conn, ch <-chan Square) {
	for {
		sq := <-ch
		msg := fmt.Sprint(strconv.Itoa(sq.row), strconv.Itoa(sq.col), strconv.Itoa(int(sq.val)))
		err := conn.WriteMessage(1, []byte(msg))
		if err != nil {
			log.Fatal("Exit")
		}
	}
}

func (sh *socketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicln("Websocket handshake failed")
	}

	ch := make(chan Square, 100)

	// Initialize & solve puzzle
	board := InitBoard()
	game := NewGame(board)
	go game.Solve(BruteForce, ch)
	go reader(conn)
	go writer(conn, ch)
}

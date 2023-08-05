package src

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/websocket"
)

// Parameterize if needed
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type socketHandler struct {
	log *log.Logger
}

func NewSocketHandler(log *log.Logger) *socketHandler {
	return &socketHandler{log: log}
}

func (sh *socketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		sh.log.Fatalln("Websocket handshake failed", err)
	}

	ch := make(chan Square, 100)
	quit := make(chan bool)

	// Reader goroutine
	// Reads board and solves puzzle
	go func() {
		var msg string
		re := regexp.MustCompile("[0-9]+")

		for {
			msgType, buf, err := conn.ReadMessage()
			if err != nil {
				conn.Close()
				panic("Error when reading")
			}

			if msgType == 1 {
				switch msg = string(buf); msg {
				case "STOP":
					quit <- true
					return
				default:
					// Read board
					// Extract numbers from message as a string
					// Translate to 2D array
					posStr := re.FindAllString(msg, -1)
					var board [9][9]uint8
					for i := 0; i < len(posStr); i++ {
						val, err := strconv.Atoi(posStr[i])
						if err != nil {
							return
						}
						board[i/9][i%9] = uint8(val)
					}

					// Initialize & solve puzzle
					game := NewGame(board)
					game.Solve(BruteForce, ch)
					game.PrintBoard()
				}
			}
		}
	}()

	// Writer goroutine
	// Extracts position and value from channel
	// Sends to client
	go func() {
		for {
			square := <-ch
			err := conn.WriteMessage(1, []byte(fmt.Sprintf("%d:%d:%d", square.row, square.col, square.val)))
			if err != nil {
				return
			}
		}
	}()
}

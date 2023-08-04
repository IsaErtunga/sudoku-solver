package main

import (
	"log"
	"os"

	"github.com/IsaErtunga/sudoku-solver/sudoku-go/src"
)

var serviceName string = "sudoku-live"
var bindAdress string = ":9090"

func main() {
	log := log.New(os.Stdout, serviceName, log.LstdFlags)
	server := src.NewServer(serviceName, bindAdress, log)
	socketHandler := src.NewSocketHandler(log)
	server.InitServer(socketHandler)
}

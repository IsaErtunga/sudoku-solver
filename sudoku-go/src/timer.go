package src

import (
	"log"
	"time"
)

func Time(id string) {
	elapsed := time.Since(time.Now())
	log.Printf("%s took %s", id, elapsed)
}

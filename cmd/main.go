package main

import (
	"log"

	"github.com/rohit21755/go_webrtc/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}

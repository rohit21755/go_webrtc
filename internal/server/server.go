package server

import (
	"flag"
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	addr = flag.String("addr", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func RUn() error {
	flag.Parse()
	app := fiber.New()
	if *addr == ":" {
		*addr = ":8000"
	}

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websockets")
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsockets))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket")
	app.Get("/stream/:ssuid/chat/websocket")
	app.Get("/stream/:ssuid/viewer/websocket")

	app.Listen(*addr)
}

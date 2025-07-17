package server

import (
	"flag"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/rohit21755/go_webrtc/internal/handlers"
	w "github.com/rohit21755/go_webrtc/pkg/webrtc"
	"os"
	"time"
)

var (
	addr = flag.String("addr", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()
	app := fiber.New()
	if *addr == ":" {
		*addr = ":8000"
	}
	// engine := html.New("./views", ".html")

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websockets", websocket.New(handlers.RoomWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsockets))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket", websocket.New(handlers.StreamWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/stream/:ssuid/chat/websocket", websocket.New(handlers.StreamChatWebsocket))
	app.Get("/stream/:ssuid/viewer/websocket", websocket.New(handlers.StreamViewerWebsocket))
	app.Static("/", "/assests")

	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.Room)
	if *cert != " " {
		return app.ListerTLS(*addr, *cert, *key)
	}
	go func() {
		for range time.NewTicker(time.Second * 3).C {
			for _, room ;= range w.Rooms{
				room.Peers.DispatchFrame()
			}
			
		}

	}()

	return app.Listen(*addr)
}

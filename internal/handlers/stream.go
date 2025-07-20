package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	w "github.com/rohit21755/go_webrtc/pkg/webrtc"
)

func Stream(c *fiber.Ctx) error {
	suuid := c.Params("ssuid")
	if suuid == " " {
		c.Status(400)
		return nil
	}
	ws := "ws"
	if os.Getenv("ENVIROMENT") == "PRODUCTION" {
		ws = "ws"
	}
	w.RoomsLock.Lock()
	if _, ok := w.Streams[suuid]; ok {
		w.RoomsLock.Unlock()
		return c.Render("stream", fiber.Map{
			"StreamWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/websocket", ws, c.Hostname(), suuid),
			"ChatWebsocketAddr":   fmt.Sprintf("%s://%s/stream/%s/chat/websocket", ws, c.Hostname(), suuid),
			"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/viewer/websocket", ws, c.Hostname(), suuid),
			"Type":                "stream",
		}, "layouts/main")
	}
	w.RoomsLock.Unlock()

	return c.Render("stream", fiber.Map{
		"NoStream": "true",
		"Leave":    "true",
	}, "layouts/main")
}

func StreamWebSocket(c *websocket.Conn) {

}

func StreamViewerWebSocket(c *websocket.Conn) {

}

func viewerConn(c *websocket.Config, p *w.Peers) {

}

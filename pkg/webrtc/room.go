package webrtc

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/rohit21755/go_webrtc/pkg/webrtc"
)

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuration

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Print(err)
		return
	}

	newPeer := PeerConnectionState{
		PeerConnection: peerConnection,
		Websocket:      &ThreadsafeWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}

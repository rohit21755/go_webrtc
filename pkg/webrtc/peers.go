package webrtc

import (
	"sync"

	"github.com/rohit21755/go_webrtc/pkg/webrtc"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalsStaticRTP
}

func (p *Peer) DispatchKeyFrame() {

}

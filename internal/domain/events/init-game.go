package events

import (
	"github.com/lccmrx/cwt/internal/domain/server"
)

func HandleInitGameEvent(server *server.ServerState, event Event) {
	server.StartMatch()
}

package events

import "github.com/lccmrx/cwt/internal/domain/server"

func HandleShutdownEvent(server *server.ServerState, _ Event) {
	server.EndMatch()
}

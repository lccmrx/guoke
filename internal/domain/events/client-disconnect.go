package events

import (
	"regexp"

	"github.com/lccmrx/cwt/internal/domain/server"
)

var (
	reClientDisconnectEvent = regexp.MustCompile(`(?P<playerid>\d+)`)
)

type ClientDisconnectEvent Event

func (event *ClientDisconnectEvent) Participant() string {
	matches := reClientDisconnectEvent.FindStringSubmatch(string(event.Data))
	playerId := string(matches[reClientDisconnectEvent.SubexpIndex("playerid")])

	return playerId
}

func HandleClientDisconnectEvent(server *server.ServerState, event Event) {
	disconnectEvent := ClientDisconnectEvent(event)
	playerId := disconnectEvent.Participant()
	delete(server.CurrentMatch.Players, playerId)
}

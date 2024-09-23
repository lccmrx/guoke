package events

import (
	"regexp"

	"github.com/lccmrx/cwt/internal/domain/server"
)

type ClientDisconnectEvent Event

func (event *ClientDisconnectEvent) Participant() string {
	re := regexp.MustCompile(`(?P<playerid>\d+)`)
	matches := re.FindStringSubmatch(string(event.Data))
	playerId := string(matches[re.SubexpIndex("playerid")])

	return playerId
}

func HandleClientDisconnectEvent(server *server.ServerState, event Event) {
	disconnectEvent := ClientDisconnectEvent(event)
	playerId := disconnectEvent.Participant()
	delete(server.CurrentMatch.Players, playerId)
}

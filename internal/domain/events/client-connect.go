package events

import (
	"regexp"

	"github.com/lccmrx/cwt/internal/domain/player"
	"github.com/lccmrx/cwt/internal/domain/server"
)

type ClientConnectEvent Event

func (event *ClientConnectEvent) Participant() string {
	re := regexp.MustCompile(`(?P<playerid>\d+)`)
	matches := re.FindStringSubmatch(string(event.Data))
	playerId := string(matches[re.SubexpIndex("playerid")])

	return playerId
}

func HandleClientConnectEvent(server *server.ServerState, event Event) {
	connectEvent := ClientConnectEvent(event)
	playerId := connectEvent.Participant()
	server.CurrentMatch.Players[playerId] = &player.Player{}
}

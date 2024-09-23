package events

import (
	"regexp"

	"github.com/lccmrx/cwt/internal/domain/player"
	"github.com/lccmrx/cwt/internal/domain/server"
)

var (
	reClientConnectEvent = regexp.MustCompile(`(?P<playerid>\d+)`)
)

type ClientConnectEvent Event

func (event *ClientConnectEvent) Participant() string {
	matches := reClientConnectEvent.FindStringSubmatch(string(event.Data))
	playerId := string(matches[reClientConnectEvent.SubexpIndex("playerid")])

	return playerId
}

func HandleClientConnectEvent(server *server.ServerState, event Event) {
	connectEvent := ClientConnectEvent(event)
	playerId := connectEvent.Participant()
	server.CurrentMatch.Players[playerId] = &player.Player{}
}

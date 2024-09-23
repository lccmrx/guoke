package events

import (
	"regexp"

	"github.com/lccmrx/cwt/internal/domain/server"
)

var (
	reClientInfoChangedEvent = regexp.MustCompile(`(?P<playerid>\d+) n\\(?P<playername>[^\\]+)\\t`)
)

type ClientUserinfoChangedEvent Event

func (event *ClientUserinfoChangedEvent) Participant() (string, string) {
	matches := reClientInfoChangedEvent.FindStringSubmatch(string(event.Data))
	playerId := string(matches[reClientInfoChangedEvent.SubexpIndex("playerid")])
	playerName := string(matches[reClientInfoChangedEvent.SubexpIndex("playername")])

	return playerId, playerName
}

func HandleClientInfoChangedEvent(server *server.ServerState, event Event) {
	userInfoChangedEvent := ClientUserinfoChangedEvent(event)
	playerId, playerName := userInfoChangedEvent.Participant()

	player := server.CurrentMatch.Players[playerId]
	player.Name = playerName
}

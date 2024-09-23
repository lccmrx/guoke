package events

import (
	"regexp"

	"github.com/lccmrx/cwt/internal/domain/server"
)

type ClientUserinfoChangedEvent Event

func (event *ClientUserinfoChangedEvent) Participant() (string, string) {
	re := regexp.MustCompile(`(?P<playerid>\d+) n\\(?P<playername>[^\\]+)\\t`)
	matches := re.FindStringSubmatch(string(event.Data))
	playerId := string(matches[re.SubexpIndex("playerid")])
	playerName := string(matches[re.SubexpIndex("playername")])

	return playerId, playerName
}

func HandleClientInfoChangedEvent(server *server.ServerState, event Event) {
	userInfoChangedEvent := ClientUserinfoChangedEvent(event)
	playerId, playerName := userInfoChangedEvent.Participant()

	player := server.CurrentMatch.Players[playerId]
	player.Name = playerName
}

package events

import (
	"regexp"
	"strconv"

	"github.com/lccmrx/cwt/internal/domain/means"
	"github.com/lccmrx/cwt/internal/domain/player"
	"github.com/lccmrx/cwt/internal/domain/server"
)

type KillEvent Event

func (event *KillEvent) Participants() (string, string, string) {
	re := regexp.MustCompile(`(?P<killer>\d+) (?P<killed>\d+) (?P<mean>\d+)`)
	matches := re.FindStringSubmatch(string(event.Data))
	killerId := string(matches[re.SubexpIndex("killer")])
	killedId := string(matches[re.SubexpIndex("killed")])
	meanId := string(matches[re.SubexpIndex("mean")])

	return killerId, killedId, meanId
}

func HandleKillEvent(server *server.ServerState, event Event) {
	killEvent := KillEvent(event)
	killerId, killedId, meanId := killEvent.Participants()

	server.CurrentMatch.TotalKills++

	mean, _ := strconv.Atoi(meanId)
	server.CurrentMatch.Means[means.MeanNameMap[means.Mean(mean)]]++

	if killerId == player.WORLD {
		server.CurrentMatch.Players[killedId].KillCount--
		return
	}

	if killerId == killedId {
		server.CurrentMatch.Players[killerId].DeathCount++
		return
	}

	server.CurrentMatch.Players[killerId].KillCount++
	server.CurrentMatch.Players[killedId].DeathCount++
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lccmrx/cwt/internal/domain/events"
	"github.com/lccmrx/cwt/internal/domain/server"
)

func main() {
	fileData, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	server := server.New()

	for _, event := range events.Events(fileData) {
		events.Handle(server, event)
	}

	type report struct {
		TotalKills  int            `json:"total_kills"`
		Players     []string       `json:"players"`
		Kills       map[string]int `json:"kills"`
		KillsByMean map[string]int `json:"kills_by_means"`
	}

	gameReports := make(map[string]report, 0)

	for i, match := range server.Matches {

		report := report{
			TotalKills:  match.TotalKills,
			Players:     make([]string, 0),
			Kills:       make(map[string]int),
			KillsByMean: make(map[string]int),
		}

		for _, player := range match.Players {
			report.Players = append(report.Players, player.Name)
			report.Kills[player.Name] = player.KillCount
		}

		report.KillsByMean = match.Means

		gameReports[fmt.Sprintf("game_%d", i+1)] = report
	}

	json2, _ := json.MarshalIndent(gameReports, "", "  ")
	fmt.Println(string(json2))
}

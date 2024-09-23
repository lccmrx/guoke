package server

import (
	"github.com/lccmrx/cwt/internal/domain/player"
)

type ServerState struct {
	CurrentMatch *Match   `json:"-"`
	Matches      []*Match `json:"-"`
}

type Match struct {
	TotalKills int
	Players    map[string]*player.Player
	Means      map[string]int
}

func New() *ServerState {
	return &ServerState{}
}

func (s *ServerState) StartMatch() {
	// if there is a current match, then we need to close it
	// delete, and start a new one
	if s.CurrentMatch != nil {
		s.Matches = s.Matches[:len(s.Matches)-1]
	}

	match := Match{Players: make(map[string]*player.Player), Means: make(map[string]int)}
	s.CurrentMatch = &match
	s.Matches = append(s.Matches, &match)
}

func (s *ServerState) EndMatch() {
	s.CurrentMatch = nil
}

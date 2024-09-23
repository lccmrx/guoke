package events

import (
	"bytes"
	"fmt"
	"iter"
	"regexp"

	"github.com/lccmrx/cwt/internal/domain/server"
)

const (
	InitGame              EventName = "InitGame"
	ShutdownGame          EventName = "ShutdownGame"
	Kill                  EventName = "Kill"
	ClientConnect         EventName = "ClientConnect"
	ClientDisconnect      EventName = "ClientDisconnect"
	ClientUserinfoChanged EventName = "ClientUserinfoChanged"
)

var (
	reEvent = regexp.MustCompile(`(?P<stopwatch>\d+:\d+) (?:-{60}|(?P<event>\w+): ?(?P<data>.*))`)
)

type EventName string

type Event struct {
	Stopwatch string
	Event     EventName
	Data      string
}

type EventHandler func(server *server.ServerState, event Event)

var eventHandlerMap = map[EventName]EventHandler{
	InitGame:              HandleInitGameEvent,
	ShutdownGame:          HandleShutdownEvent,
	ClientConnect:         HandleClientConnectEvent,
	ClientUserinfoChanged: HandleClientInfoChangedEvent,
	Kill:                  HandleKillEvent,
	// ClientDisconnect:      HandleClientDisconnectEvent,
	// removed due to unkwnown game rule for handling the score when a user disconnects
}

func Handle(server *server.ServerState, event Event) error {
	handler, ok := eventHandlerMap[event.Event]
	if !ok {
		return fmt.Errorf("no handler for event %s", event.Event)
	}

	handler(server, event)
	return nil
}

func Events(data []byte) iter.Seq2[int, Event] {
	return func(yield func(i int, data Event) bool) {
		for len(data) > 0 {
			line, rest, _ := bytes.Cut(data, []byte{'\n'})

			matches := reEvent.FindStringSubmatch(string(line))
			stopawatch := string(matches[reEvent.SubexpIndex("stopwatch")])
			eventName := EventName(matches[reEvent.SubexpIndex("event")])
			eventData := string(matches[reEvent.SubexpIndex("data")])

			event := Event{
				Stopwatch: stopawatch,
				Event:     eventName,
				Data:      eventData,
			}

			if !yield(0, event) {
				return
			}
			data = rest
		}
	}
}

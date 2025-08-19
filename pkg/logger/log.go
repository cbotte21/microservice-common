package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type CRUD = string

const (
	Create   CRUD = "Create"
	Retrieve CRUD = "Retrieve"
	Update   CRUD = "Update"
	Delete   CRUD = "Delete"
)

type Event struct {
	Timestamp time.Time `json:"timestamp"`
	Operation CRUD      `json:"operation"`
	Location  string    `json:"location"`
	Context   any       `json:"context,omitempty"`
}

type Logger interface {
	Run()

	Write(event Event)
	Err() <-chan error
}

type STDOUTLogger struct {
	events chan<- Event
	errors <-chan error

	eventsQueueLength uint8
}

func NewSTDOUTLogger() Logger {
	var eventsQueueLength uint8

	if v, err := strconv.Atoi(os.Getenv("LOGGER_QUEUE_LENGTH")); err != nil {
		eventsQueueLength = uint8(v)
	} else {
		_ = fmt.Errorf("invalid logger queue length")
	}

	return &STDOUTLogger{eventsQueueLength: eventsQueueLength}
}

func (l *STDOUTLogger) Run() {
	events := make(chan Event, l.eventsQueueLength)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	for ev := range events {
		b, err := json.Marshal(ev)
		if err != nil {
			errors <- err
			continue
		}
		fmt.Println(string(b))
	}
}

func (l *STDOUTLogger) Write(event Event) {
	l.events <- event
}

func (l *STDOUTLogger) Err() <-chan error {
	return l.errors
}

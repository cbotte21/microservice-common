package logger

import (
	"encoding/json"
	"fmt"
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
	Write(event Event)
	Err() <-chan error
}

type TransactionLogger struct {
	events chan<- Event
	errors <-chan error
}

func NewLogger(queueLength uint8) Logger {
	l := &TransactionLogger{}
	go l.run(queueLength)
	return l
}

func (l *TransactionLogger) run(queueLength uint8) {
	events := make(chan Event, queueLength)
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

func (l *TransactionLogger) Write(event Event) {
	l.events <- event
}

func (l *TransactionLogger) Err() <-chan error {
	return l.errors
}

package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type CRUD = string

const (
	Create   CRUD = "Create"
	Retrieve CRUD = "Retrievee"
	Update   CRUD = "Update"
	Delete   CRUD = "Delete"
)

type Event struct {
	Timestamp time.Time `json:"timestamp"`
	Operation CRUD      `json:"operation"`
	Target    string    `json:"target"`
	Context   any       `json:"context,omitempty"`
}

type Logger interface {
	Write(event Event)
}

type StandardLogger struct {
	ctx    context.Context
	cancel context.CancelFunc
	logs   chan Event
	errors chan<- error
}

func NewLogger(ctx context.Context, errors chan<- error, queueLength uint8) Logger {
	if queueLength <= 0 {
		queueLength = 10
	}
	cctx, cancel := context.WithCancel(ctx)
	l := &StandardLogger{
		ctx:    cctx,
		cancel: cancel,
		logs:   make(chan Event, queueLength),
		errors: errors,
	}
	go l.run()
	return l
}

func (l *StandardLogger) run() {
	for {
		select {
		case <-l.ctx.Done():
			// Finish queue?
			fmt.Println("Terminating log queue...")
			close(l.logs)
			return
		case ev := <-l.logs:
			b, err := json.Marshal(ev)
			if err != nil {
				l.errors <- err
				continue
			}
			fmt.Println(string(b))
		}
	}
}

func (l *StandardLogger) Write(event Event) {
	l.logs <- event
}

package logger

import (
	"context"
	"math"
	"testing"
	"time"

	"github.com/cbotte21/microservice-common/pkg/logger"
)

var queueLength uint8
var errorsChan chan error
var l logger.Logger

func init() {
	queueLength = 4
	errorsChan = make(chan error, 1)
	l = logger.NewLogger(context.Background(), errorsChan, queueLength)
}

func TestLoggingQueue(t *testing.T) {
	for i := range int(math.Pow(float64(queueLength), 3)) {
		l.Write(logger.Event{
			Operation: logger.Delete,
			Target:    "queue test",
			Timestamp: time.Now(),
			Context:   map[string]any{"index": i},
		})
	}
}

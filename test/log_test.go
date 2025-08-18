package logger

import (
	"math"
	"testing"
	"time"

	"github.com/cbotte21/microservice-common/pkg/logger"
)

var queueLength uint8
var l logger.Logger

func init() {
	queueLength = 16
	l = logger.NewLogger(queueLength)
}

func TestLoggingQueue(t *testing.T) {
	for i := range int(math.Pow(float64(queueLength), 3)) {
		l.Write(logger.Event{
			Operation: logger.Delete,
			Location:  "queue test",
			Timestamp: time.Now(),
			Context:   map[string]any{"index": i},
		})
	}
}

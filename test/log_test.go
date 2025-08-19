package logger

import (
	"math"
	"testing"
	"time"

	"github.com/cbotte21/microservice-common/pkg/logger"
)

func TestLoggingQueue(t *testing.T) {
	// Set an arbitrary queue length for this test and have it automatically
	// restored after the test finishes.
	t.Setenv("LOGGER_QUEUE_LENGTH", "8")

	l := logger.NewSTDOUTLogger()

	// Start the logger; this goroutine will exit when the process exits.
	// (The logger currently doesn't expose a shutdown, so we just let it run.)
	go l.Run()
	time.Sleep(50 * time.Millisecond)

	// Write enough events to exercise the queueing behavior.
	// Using 8^3 = 512 events which is several times the buffer length.
	for i := 0; i < int(math.Pow(8, 3)); i++ {
		l.Write(logger.Event{
			Operation: logger.Delete,
			Location:  "queue test",
			Timestamp: time.Now(),
			Context:   map[string]any{"index": i},
		})
	}

	// Briefly wait to allow the background goroutine to drain the queue.
	time.Sleep(50 * time.Millisecond)
}

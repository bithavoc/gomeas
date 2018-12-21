package gomeas

import (
	"testing"
	"time"
)

func TestStopwatchTest(t *testing.T) {
	sw := StartStopwatch("2 second call")
	time.Sleep(2 * time.Second)
	sw.Checkpoint("less than a second")
	time.Sleep(800 * time.Millisecond)
	sw.Checkpoint("1 second call")
	time.Sleep(1 * time.Second)
	measurements := sw.Complete()
	measurementCount := len(measurements)
	if measurementCount != 3 {
		t.Fatalf("measurements count does not match (count %d)", measurementCount)
	}
	m := measurements[0]
	if m.Duration.Seconds() < 2 {
		t.Fatalf("measurement is not accurate (duration %s)", m.Duration)
	}
	if m.Name != "2 second call" {
		t.Fatalf("measurement name does not match (name %s)", m.Name)
	}
	m = measurements[1]
	if m.Duration.Seconds() > 1 {
		t.Fatalf("measurement is not accurate (duration %s)", m.Duration)
	}
	if m.Name != "less than a second" {
		t.Fatalf("measurement name does not match (name %s)", m.Name)
	}
	m = measurements[2]
	if m.Duration.Seconds() < 1 {
		t.Fatalf("measurement is not accurate (duration %s)", m.Duration)
	}
	if m.Name != "1 second call" {
		t.Fatalf("measurement name does not match (name %s)", m.Name)
	}
}

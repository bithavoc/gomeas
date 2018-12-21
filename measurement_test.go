package gomeas

import (
	"testing"
	"time"
)

func TestMeasurementTest(t *testing.T) {
	m := StartMeasurement("2 second call")
	time.Sleep(2 * time.Second)
	m.Complete()

	if m.Duration.Seconds() < 2 {
		t.Fatalf("measurement is not accurate (duration %s)", m.Duration)
	}
	if m.Name != "2 second call" {
		t.Fatalf("measurement name does not match (name %s)", m.Name)
	}
}

func TestCompletedMeasurementsTest(t *testing.T) {
	m := CompletedMeasurements{
		Measurement{
			Name:     "short",
			Duration: 10 * time.Second,
		},
		Measurement{
			Name:     "long",
			Duration: 5 * time.Minute,
		},
	}
	if m.String() != "short: 10s, long: 5m0s" {
		t.Fatalf("measurements description mismatch(duration %s)", m.String())
	}
}

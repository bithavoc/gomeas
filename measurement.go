package gomeas

import (
	"fmt"
	"strings"
	"time"
)

// Measurement contains the measurements data for a time period
type Measurement struct {
	Name        string
	StartedAt   time.Time
	CompletedAt time.Time
	Duration    time.Duration
}

// StartMeasurement returns a measurement started at time.Now()
func StartMeasurement(name string) Measurement {
	return Measurement{
		Name:      name,
		StartedAt: time.Now(),
	}
}

// Complete completes the measurement
func (m *Measurement) Complete() {
	m.CompletedAt = time.Now()
	m.Duration = m.CompletedAt.Sub(m.StartedAt)
}

func (m *Measurement) String() string {
	return fmt.Sprintf("%s: %s", m.Name, m.Duration)
}

// CompletedMeasurements contains a slice of completed measurements
type CompletedMeasurements []Measurement

// String returns an string representation of all measures
func (m CompletedMeasurements) String() string {
	ms := make([]string, len(m))
	for i, m := range m {
		ms[i] = m.String()
	}
	return strings.Join(ms, ", ")
}

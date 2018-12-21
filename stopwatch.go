package gomeas

// Stopwatch measures multiple checkpoints and report them all at once
type Stopwatch struct {
	measurements CompletedMeasurements
	current      Measurement
}

// StartStopwatch starts a stopwatch with the given initial checkpoint
func StartStopwatch(initialCheckpoint string) Stopwatch {
	w := Stopwatch{}
	w.Checkpoint(initialCheckpoint)
	return w
}

// Checkpoint marks the end of the current timer and starts another one
func (sw *Stopwatch) Checkpoint(name string) {
	sw.completeCurrent()
	sw.current = StartMeasurement(name)
}

// Complete completes the last measurement of the stop watch and returns the completed measurements
func (sw *Stopwatch) Complete() CompletedMeasurements {
	sw.completeCurrent()
	return sw.measurements
}

func (sw *Stopwatch) completeCurrent() {
	if sw.current == (Measurement{}) {
		return
	}
	sw.current.Complete()
	sw.measurements = append(sw.measurements, sw.current)
	sw.current = Measurement{}
}

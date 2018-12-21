## gomeas
[![CircleCI](https://circleci.com/gh/bithavoc/gomeas.svg?style=svg&circle-token=7bc0eb45018b5021c738e6bd0c40f16fec53f879)](https://circleci.com/gh/bithavoc/gomeas)
[![GoCover](https://gocover.io/_badge/github.com/bithavoc/gomeas)](https://gocover.io/github.com/bithavoc/gomeas)
[![Go Report Card](https://goreportcard.com/badge/github.com/bithavoc/gomeas)](https://goreportcard.com/report/github.com/bithavoc/gomeas)

> Time measurements for the lazy

### Example:

```go
  sw := StartStopwatch("2 second call")
  time.Sleep(2 * time.Second)
  sw.Checkpoint("less than a second")
  time.Sleep(800 * time.Millisecond)
  sw.Checkpoint("1 second call")
  time.Sleep(1 * time.Second)
  measurements := sw.Complete()

  fmt.Println(measurements.String())
  // "2 second call: 2s, less than a second: 802ms, 1 second call": 1s"
```

### Test

```
go test
```

with coverage:

```
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```


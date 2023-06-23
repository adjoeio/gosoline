package producer

import "time"

type CircuitBreakerSettings struct {
	MaxFailures int64         `cfg:"max_failures" default:"10"`
	RetryDelay  time.Duration `cfg:"retry_delay" default:"1m"`
}

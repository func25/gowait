package gowait

import "time"

type WaitOptGen struct{}
type waitOpt func(*waitConfig)

type waitConfig struct {
	panicRetry         bool
	panicRetryDuration time.Duration
}

func (r *waitConfig) applyOpts(opts ...waitOpt) *waitConfig {
	*r = waitConfig{
		panicRetry:         true,
		panicRetryDuration: 15 * time.Second,
	}

	for _, v := range opts {
		v(r)
	}

	return r
}

func (o WaitOptGen) PanicRetry(retry bool, duration time.Duration) waitOpt {
	return func(r *waitConfig) {
		r.panicRetry = retry
		r.panicRetryDuration = duration
	}
}

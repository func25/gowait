package gowait

import "time"

type RepeatOptGen struct{}
type repeatOpt func(*repeatConfig)

type repeatConfig struct {
	zeroDuration time.Duration
	minDuration  time.Duration
	waitConfig
}

func (r *repeatConfig) init() *repeatConfig {
	*r = repeatConfig{
		zeroDuration: time.Second,
		minDuration:  time.Second,
	}
	r.waitConfig.init()

	return r
}

func (r *repeatConfig) applyOpts(opts ...repeatOpt) *repeatConfig {
	r.init()

	for _, v := range opts {
		v(r)
	}

	return r
}

func (r repeatConfig) applyDuration(t time.Duration) (n time.Duration) {
	n = t
	if r.zeroDuration > 0 && t <= 0 {
		n = r.zeroDuration
	}

	if n < r.minDuration {
		n = r.minDuration
	}

	return
}

func (o RepeatOptGen) ZeroDuration(zero time.Duration) repeatOpt {
	return func(r *repeatConfig) {
		r.zeroDuration = zero
	}
}

func (o RepeatOptGen) MinDuration(min time.Duration) repeatOpt {
	return func(r *repeatConfig) {
		r.minDuration = min
	}
}

func (o RepeatOptGen) PanicRetry(retry bool, duration time.Duration) repeatOpt {
	return func(r *repeatConfig) {
		r.panicRetry = retry
		r.panicRetryDuration = duration
	}
}

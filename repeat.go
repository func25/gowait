package gowait

import (
	"time"
)

func ScheduleFuncLoop(t time.Time, f func() *time.Time, opts ...repeatOpt) *time.Timer {
	d := t.Sub(time.Now())
	cfg := (&repeatConfig{}).applyOpts(opts...)
	return scheduleFuncLoop(d, f, *cfg)
}

// hidden from external
func scheduleFuncLoop(t time.Duration, f func() *time.Time, r repeatConfig) *time.Timer {
	return durationFunc(t, repeatScheduleFunc(f, r), r.waitConfig)
}

func repeatScheduleFunc(f func() *time.Time, r repeatConfig) func() {
	return func() {
		nextTime := f()
		if nextTime != nil {
			interval := r.applyDuration(nextTime.Sub(time.Now()))
			scheduleFuncLoop(interval, f, r)
		}
	}
}

func DurationFuncLoop(duration time.Duration, f func() *time.Duration, opts ...repeatOpt) *time.Timer {
	cfg := (&repeatConfig{}).applyOpts(opts...)
	return durationFuncLoop(duration, f, *cfg)
}

// hidden from external
func durationFuncLoop(duration time.Duration, f func() *time.Duration, r repeatConfig) *time.Timer {
	return durationFunc(duration, repeatDurationFunc(f, r), r.waitConfig)
}

func repeatDurationFunc(f func() *time.Duration, r repeatConfig) func() {
	return func() {
		nextTime := f()
		if nextTime != nil {
			*nextTime = r.applyDuration(*nextTime)
			durationFuncLoop(*nextTime, f, r)
		}
	}
}

package gowait

import (
	"time"
)

func ScheduleFuncLoop(t time.Time, job func() *time.Time, opts ...repeatOpt) *time.Timer {
	d := t.Sub(time.Now())
	cfg := (&repeatConfig{}).applyOpts(opts...)
	return scheduleJobLoop(d, job, *cfg)
}

// hidden from external
func scheduleJobLoop(t time.Duration, job func() *time.Time, cfg repeatConfig) *time.Timer {
	return durationJob(t, repeatScheduleJob(job, cfg), cfg.waitConfig)
}

func repeatScheduleJob(job func() *time.Time, cfg repeatConfig) func() {
	return func() {
		nextTime := job()
		if nextTime != nil {
			interval := cfg.applyDuration(nextTime.Sub(time.Now()))
			scheduleJobLoop(interval, job, cfg)
		}
	}
}

func DurationJobLoop(duration time.Duration, job func() *time.Duration, opts ...repeatOpt) *time.Timer {
	cfg := (&repeatConfig{}).applyOpts(opts...)
	return durationJobLoop(duration, job, *cfg)
}

// hidden from external
func durationJobLoop(duration time.Duration, job func() *time.Duration, r repeatConfig) *time.Timer {
	return durationJob(duration, repeatDurationFunc(job, r), r.waitConfig)
}

func repeatDurationFunc(job func() *time.Duration, r repeatConfig) func() {
	return func() {
		nextTime := job()
		if nextTime != nil {
			*nextTime = r.applyDuration(*nextTime)
			durationJobLoop(*nextTime, job, r)
		}
	}
}

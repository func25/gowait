package gowait

import (
	"time"
)

func ScheduleJobLoop(job func() *time.Time, t time.Time, opts ...repeatOpt) *time.Timer {
	d := t.Sub(time.Now())
	cfg := (&repeatConfig{}).applyOpts(opts...)
	return scheduleJobLoop(job, d, *cfg)
}

// hidden from external
func scheduleJobLoop(job func() *time.Time, t time.Duration, cfg repeatConfig) *time.Timer {
	return durationJob(repeatScheduleJob(job, cfg), t, cfg.waitConfig)
}

func repeatScheduleJob(job func() *time.Time, cfg repeatConfig) func() {
	return func() {
		nextTime := job()
		if nextTime != nil {
			interval := cfg.applyDuration(nextTime.Sub(time.Now()))
			scheduleJobLoop(job, interval, cfg)
		}
	}
}

func DurationJobLoop(job func() *time.Duration, duration time.Duration, opts ...repeatOpt) *time.Timer {
	cfg := (&repeatConfig{}).applyOpts(opts...)
	return durationJobLoop(job, duration, *cfg)
}

// hidden from external
func durationJobLoop(job func() *time.Duration, duration time.Duration, r repeatConfig) *time.Timer {
	return durationJob(repeatDurationFunc(job, r), duration, r.waitConfig)
}

func repeatDurationFunc(job func() *time.Duration, r repeatConfig) func() {
	return func() {
		nextTime := job()
		if nextTime != nil {
			*nextTime = r.applyDuration(*nextTime)
			durationJobLoop(job, *nextTime, r)
		}
	}
}

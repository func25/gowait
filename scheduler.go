package gowait

import (
	"time"
)

// return controller to stop the function
func DurationJob(job func(), duration time.Duration, opts ...waitOpt) *time.Timer {
	cfg := (&waitConfig{}).applyOpts(opts...)
	return durationJob(job, duration, *cfg)
}

func durationJob(job func(), duration time.Duration, cfg waitConfig) *time.Timer {
	if duration <= 0 {
		go recoverJob(job, cfg)()
		return nil
	}
	return time.AfterFunc(duration, recoverJob(job, cfg))
}

// return controller to stop the function
func ScheduleJob(job func(), t time.Time, opts ...waitOpt) *time.Timer {
	d := t.Sub(time.Now())
	return DurationJob(job, d, opts...)
}

func recoverJob(job func(), cfg waitConfig) func() {
	return func() {
		defer func() {
			r := recover()
			if r != nil {
				if cfg.panicRetry {
					durationJob(job, cfg.panicRetryDuration, cfg)
				}
			}
		}()

		job()
	}
}

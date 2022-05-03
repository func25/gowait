package gowait

import (
	"time"
)

// return controller to stop the function
func DurationJob(duration time.Duration, job func(), opts ...waitOpt) *time.Timer {
	cfg := (&waitConfig{}).applyOpts(opts...)
	return durationJob(duration, job, *cfg)
}

func durationJob(duration time.Duration, job func(), cfg waitConfig) *time.Timer {
	if duration <= 0 {
		go recoverJob(cfg, job)()
		return nil
	}
	return time.AfterFunc(duration, recoverJob(cfg, job))
}

// return controller to stop the function
func ScheduleJob(t time.Time, job func(), opts ...waitOpt) *time.Timer {
	d := t.Sub(time.Now())
	return DurationJob(d, job, opts...)
}

func recoverJob(cfg waitConfig, job func()) func() {
	return func() {
		defer func() {
			r := recover()
			if r != nil {
				if cfg.panicRetry {
					durationJob(cfg.panicRetryDuration, job, cfg)
				}
			}
		}()

		job()
	}
}

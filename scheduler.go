package gowait

import (
	"time"
)

// return controller to stop the function
func DurationFunc(duration time.Duration, f func(), opts ...waitOpt) *time.Timer {
	cfg := (&waitConfig{}).applyOpts(opts...)
	return durationFunc(duration, f, *cfg)
}

func durationFunc(duration time.Duration, f func(), cfg waitConfig) *time.Timer {
	if duration < 0 {
		go recoverFunc(cfg, f)()
		return nil
	}
	return time.AfterFunc(duration, recoverFunc(cfg, f))
}

// return controller to stop the function
func ScheduleFunc(t time.Time, f func(), opts ...waitOpt) *time.Timer {
	d := t.Sub(time.Now())
	return DurationFunc(d, f, opts...)
}

func recoverFunc(cfg waitConfig, f func()) func() {
	return func() {
		defer func() {
			r := recover()
			if r != nil {
				if cfg.panicRetry {
					DurationFunc(cfg.panicRetryDuration, f)
				}
			}
		}()

		f()
	}
}

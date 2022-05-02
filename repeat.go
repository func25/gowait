package gowait

import (
	"time"
)

type RepeatOpt func(*time.Duration)

func OptMinDuration(min time.Duration) func(*time.Duration) {
	return func(t *time.Duration) {
		if *t < 0 {
			*t = min
		}
	}
}

func ScheduleFuncLoop(t time.Time, f func() *time.Time, opts ...RepeatOpt) *time.Timer {
	d := t.Sub(time.Now())
	return scheduleRepeatFunc(d, f, opts...)
}

func scheduleRepeatFunc(t time.Duration, f func() *time.Time, opts ...RepeatOpt) *time.Timer {
	return DurationFunc(t, repeatScheduleFunc(f, opts...))
}

func repeatScheduleFunc(f func() *time.Time, opts ...RepeatOpt) func() {
	return func() {
		nextTime := f()
		interval := nextTime.Sub(time.Now())
		if nextTime != nil {
			for _, v := range opts {
				v(&interval)
			}
			scheduleRepeatFunc(interval, f, opts...)
		}
	}
}

func DurationFuncLoop(duration time.Duration, f func() *time.Duration, opts ...RepeatOpt) *time.Timer {
	return DurationFunc(duration, repeatDurationFunc(f, opts...))
}

func repeatDurationFunc(f func() *time.Duration, opts ...RepeatOpt) func() {
	return func() {
		nextTime := f()
		if nextTime != nil {
			for _, v := range opts {
				v(nextTime)
			}
			DurationFuncLoop(*nextTime, f, opts...)
		}
	}
}

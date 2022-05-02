package gowait

import (
	"time"
)

// return controller to stop the function
func DurationFunc(duration time.Duration, f func()) *time.Timer {
	if duration < 0 {
		go f()
		return nil
	}
	return time.AfterFunc(duration, f)
}

// return controller to stop the function
func ScheduleFunc(t time.Time, f func()) *time.Timer {
	d := t.Sub(time.Now())
	return DurationFunc(d, f)
}

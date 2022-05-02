package gowaitopts

// import "time"

// type repeatConfig struct {
// 	interval    time.Duration
// 	minDuration time.Duration
// 	repeat      bool
// }

// type RepeatOpt func(*repeatConfig)

// func OptMinDuration(min time.Duration) RepeatOpt {
// 	if min <= 0 {
// 		min = time.Second
// 	}

// 	return func(c *repeatConfig) {
// 		c.minDuration = min
// 	}
// }

// func (r *repeatConfig) applyOpts(opts ...RepeatOpt) {
// 	for _, v := range opts {
// 		v(r)
// 	}
// }

// func (r *repeatConfig) ValidateTime(t time.Time) {
// 	if r.minDuration != 0 {
// 		dis := t.Sub(time.Now())
// 		if dis <= 0 {
// 			r.interval = r.minDuration
// 		} else {
// 			r.interval = dis
// 		}
// 	}
// }

// func (r *repeatConfig) Validate(t time.Duration) {
// 	if r.minDuration > 0 {
// 		if t <= 0 {
// 			r.interval = r.minDuration
// 		} else {
// 			r.interval = t
// 		}
// 	}
// }

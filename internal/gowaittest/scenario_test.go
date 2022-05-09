package gopertest

import (
	"fmt"
	"testing"
	"time"

	"github.com/func25/gowait"
)

var startTime int64

func TestDuration(t *testing.T) {
	startTime = time.Now().Unix()
	gowait.DurationJobLoop(loopDuration, time.Second, gowait.RepeatOptGen{}.ZeroDuration(time.Second*3))
	time.Sleep(time.Hour)
}

func loopDuration() *time.Duration {
	dis := time.Now().Unix() - startTime
	fmt.Println(dis)
	x := time.Second

	if dis == 5 {
		x = 0
	}

	return &x
}

func TestTime(t *testing.T) {
	startTime = time.Now().Unix()
	g := gowait.RepeatOptGen{}
	gowait.ScheduleJobLoop(loopTime, time.Now().Add(time.Second),
		g.ZeroDuration(time.Second),
		g.MinDuration(time.Second*2),
		g.PanicRetry(true, 3*time.Second),
	)
	time.Sleep(time.Hour)
}

func loopTime() *time.Time {
	dis := time.Now().Unix() - startTime
	fmt.Println("show time:", dis)
	x := time.Now().Add(time.Second * 1) // run next 1 second

	if dis == 5 {
		fmt.Println("zeroDuration")
		x = time.Now() // test zeroDuration + minDuration
	}

	if dis > 10 { // test panic
		fmt.Println("panic")
		panic("dis > 10")
	}

	return &x
}

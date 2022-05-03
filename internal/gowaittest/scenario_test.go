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
	gowait.DurationFuncLoop(time.Second, loopDuration, gowait.RepeatOptGen{}.ZeroDuration(time.Second*3))
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
	gowait.ScheduleFuncLoop(time.Now().Add(time.Second), loopTime,
		g.ZeroDuration(time.Second*1),
		g.MinDuration(time.Second*2),
		g.PanicRetry(true, time.Second),
	)
	time.Sleep(time.Hour)
}

func loopTime() *time.Time {
	dis := time.Now().Unix() - startTime
	fmt.Println(dis)
	x := time.Now().Add(time.Second * 1)

	if dis == 5 {
		x = time.Now() // time.Now().Add(time.Second)
	}

	if dis > 10 {
		panic("dis > 10")
	}

	return &x
}

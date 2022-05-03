package gopertest

import (
	"fmt"
	"testing"
	"time"

	"github.com/func25/gowait"
)

var start int64

func TestDur(t *testing.T) {
	start = time.Now().Unix()
	gowait.DurationJob(showTime, time.Second*3)

	time.Sleep(time.Second * 4)
}

func showTime() {
	fmt.Println("show time:", time.Now().Unix()-start)
}

func TestDurRepeat(t *testing.T) {
	start = time.Now().Unix()
	gowait.DurationJobLoop(showTimeLoop, time.Second*3)

	time.Sleep(time.Second * 20)
}

func showTimeLoop() *time.Duration {
	fmt.Println("show time:", time.Now().Unix()-start)
	next := time.Second * 3

	if time.Now().Unix()-start > 10 {
		return nil
	}

	return &next
}

# gowait

Instead of cronjob, you can schedule a job run at a specific time, run after X duration, panic retry,... 

* [Installation](#installation)
* [Samples](#samples)
  * [Schedule a job](#schedule-a-job)
  * [Schedule a REPEAT job](#schedule-a-repeat-job)
  * [Options](#options-panic-custom-duration)
* [Status](#status-pre-release)

## Installation

`go get github.com/func25/gowait`

## Samples

### Schedule a job

Schedule a job run in 3 second later:

```go
var start int64

func main() {
	start = time.Now().Unix()
	gowait.DurationJob(time.Second*3, showTime)

	time.Sleep(time.Second * 4)
}

func showTime() {
	fmt.Println("show time:", time.Now().Unix() - start)
}

```

The result of course will be:
```
show time: 3
```

### Schedule a REPEAT job

Schedule a job run in 3 seconds later and also run EVERY 3 seconds, if you want to stop, just return nil on the job (showTimeLoop job for example)
```go
var start int64

func main() {
	start = time.Now().Unix()
	gowait.DurationJobLoop(showTimeLoop, time.Second*3)

	time.Sleep(time.Second * 10)
}

func showTimeLoop() *time.Duration {
	fmt.Println("show time:", time.Now().Unix()-start)
	next := time.Second * 3
	return &next
}
```

```go
show time: 3
show time: 6
show time: 9
```

### Options: panic, custom duration,...
To use our option, you should create an "option generator", in the below example, the job will run at 1 second later
- **ZeroDuration**: if the job return <= 0 duration time, then we apply 1 second to the duration (avoid spamming).
- **MinDuration**: this will be more prioritized than zeroDuration.
- **PanicRetry**: retry the job if any panic occurs and what time is it run next time.

```go
var startTime int64
func main() {
	startTime = time.Now().Unix()
	g := gowait.RepeatOptGen{}
	gowait.ScheduleFuncLoop(loopTime, time.Now().Add(time.Second),
		g.ZeroDuration(time.Second), // zeroDuration will be 2s (minDuration have higher priority)
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
		x = time.Now()
	}

	if dis > 10 {
		panic("dis > 10")
	}

	return &x
}
```

```go
show time: 1
show time: 3
show time: 5
zeroDuration
show time: 7
show time: 9
show time: 11
panic
show time: 14
panic
show time: 17
panic
show time: 20
panic
```

## Status: pre-release
This lib is under developing, please notice when using it

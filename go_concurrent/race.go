package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RaceDemo() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

// go build -race race.go
// go run -race race.go
// go test -race race.go
func main(){
	pause := func(){
		var str string
		fmt.Println("")
		fmt.Print("请按任意键继续...")
		fmt.Scanln(&str)
		fmt.Print("程序退出...")
	}

	ChanDemo() // 通道
	RaceDemo() // 竞态检测
	GoRoutineDemo() // 协程通信
	WaitGroupDemo() // 等待执行组
	LockDemo() // 锁

	pause()
}

func GoRoutineDemo(){
	//GoRoutineFailedDemo1() -- 报错
	//GoRoutineFailedDemo2() -- 不报错,但是少数据
	GoRoutineSuccessDemo()
}

func WaitGroupDemo(){
	WaitGroupDemo_chan()
	WaitGroupDemo_waitGroup()
}

func LockDemo(){
	LockDemo_Mutex()
	LockDemo_RWMutex()
}
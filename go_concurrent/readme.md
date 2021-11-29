# go_concurrent

## 1.锁
    互斥锁 mutex
    读写互斥锁 rwmutex
## 2.延迟初始化
    Once
## 3.竞态检测
    go build -race race.go 
## 4.协程/管道
    channel
    goroutine
    mutex
    
    var ch2 chan<- elementType // ch2是单向channel，只用于写数据
    var ch3 <-chan elementType // ch3是单向channel，只用于读数据
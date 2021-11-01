package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

//获取int64 获取毫秒
func Tick64(t ...time.Time) int64 {
	if len(t) == 0 {
		return time.Now().UnixNano() / 1e6
	} else {
		return t[0].UnixNano() / 1e6
	}
}

type H struct {
	m map[string]interface{}
}

func main() {
	pprof()
	for {
		time.Sleep(time.Duration(time.Second * 2))
		m := make(map[string]interface{})
		m["test"] = "test"
		h := &H{
			m: m,
		}
		fmt.Println(Tick64(), h)
	}
}

// 性能监控
func pprof() {
	fmt.Println("pprof start")
	go func() {
		// 环境:安装graphviz:https://graphviz.gitlab.io/_pages/Download/Download_windows.html
		// http监控:go tool pprof -http="localhost:8081" http://localhost:10000/debug/pprof/profile
		// file监控:go tool pprof http://localhost:10000/debug/pprof/profile
		// 		查看保存的文件: C:\Users\Administrator\pprof\pprof.samples.cpu.001.pb.gz
		// 		性能调优 go tool pprof C:\Users\Administrator\pprof\pprof.samples.cpu.001.pb.gz
		// 		file监控-ui:web
		log.Println(http.ListenAndServe(":10000", nil))
	}()
}

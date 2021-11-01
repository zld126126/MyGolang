
- [如何使用pprof](#如何使用pprof)
  - [1.引用pprof](#1引用pprof)
  - [2.启动一个runtime](#2启动一个runtime)
  - [3.启动main && 执行如下命令:](#3启动main--执行如下命令)
  - [4.安装graphviz](#4安装graphviz)
  - [5.查看保存的pprof文件 && 调优:](#5查看保存的pprof文件--调优)

## 如何使用pprof
### 1.引用pprof
    _ "net/http/pprof"
### 2.启动一个runtime
    http.ListenAndServe(":10000", nil)
### 3.启动main && 执行如下命令:
    http监控:
        go tool pprof -http="localhost:8081" http://localhost:10000/debug/pprof/profile
        浏览器访问:
            http://localhost:8081/ui/

    file监控:
        go tool pprof http://localhost:10000/debug/pprof/profile
        (pprof):web
### 4.安装graphviz
    https://graphviz.gitlab.io/_pages/Download/Download_windows.html
### 5.查看保存的pprof文件 && 调优:
    文件位置:
        C:\Users\Administrator\pprof\pprof.samples.cpu.001.pb.gz
    调优命令:
        go tool pprof C:\Users\Administrator\pprof\pprof.samples.cpu.001.pb.gz
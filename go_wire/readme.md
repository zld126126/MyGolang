步骤1：
go get github.com/google/wire

步骤2：
创建service目录下的文件 和 wire.go
****(wire.go一定要有// +build wireinject)****

步骤3：
terminal/命令行工具在包下运行命令,wire
生成wire_gen.go

步骤4：运行main.go

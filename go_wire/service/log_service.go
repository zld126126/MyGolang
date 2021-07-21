package service

import "github.com/google/wire"

type LogServiceInterface interface {
	Write()
}

type LogService struct {
	LogName string
}

func (p *LogService) Write() {
	println(p.LogName, "LogService Write")
}

func NewLogService(name string) *LogService {
	return &LogService{LogName: name}
}

var LogSet = wire.NewSet(NewLogService, wire.Bind(new(LogServiceInterface), new(*LogService)))

package boot

import (
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"

	"go_micro/pkg/controller"
	"go_micro/pkg/service"
	proto "go_micro/proto"
)

func initMicroService() micro.Service {
	// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
	service := micro.NewService(
		micro.Address(":9091"),
		micro.Name("landon.assist.service"),
		micro.Version("latest"),
		micro.AfterStart(func() error {
			fmt.Println("landon micro assist service start")
			return nil
		}),
	)
	return service
}

func initMicroWebService(microService micro.Service) web.Service {
	return web.NewService(
		web.Address(":9090"),
		web.MicroService(microService),
		web.Name("landon.web.service"),
		web.Version("latest"),
		web.AfterStart(func() error {
			fmt.Println("landon micro web service start")
			return nil
		}),
	)
}

func initMicroServiceClient(microService micro.Service) client.Client {
	return microService.Client()
}

func initAssistService(microService micro.Service) proto.QueryService {
	assist := proto.NewQueryService("landon.assist.service", microService.Client())
	return assist
}

type Handle struct {
	Web           web.Service
	Micro         micro.Service
	Client        client.Client
	AssistService proto.QueryService
	Config        *Config
	DB            *DB
}

func (p *Handle) Run() {
	// 启动服务
	go func() {
		p.Web.Handle("/", controller.InterceptController(p.AssistService, p.DB.GetGorm()))
		if err := p.Web.Run(); err != nil {
			fmt.Println(err)
		}
	}()

	// 初始化服务
	p.Micro.Init()
	// 注册服务
	proto.RegisterQueryHandler(p.Micro.Server(), new(service.AssistService))
	// 启动服务
	if err := p.Micro.Run(); err != nil {
		fmt.Println(err)
	}
}

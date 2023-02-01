package ziface

//定义一个服务器接口
type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve()
	//路由功能：给当前的服务注册一个路由方法，供客户端的链接处理使用
	AddRouter(uint32, IRouter)
	//获取当前server 的连接管理器
	GetConnMgr() IConnManager
	//注册OnConnStart 钩子函数的方法
	SetOnConnStart(func(conneciton IConneciton))
	//注册OnConnStop钩子函数的方法
	SetOnConnStop(func(conneciton IConneciton))
	//调用OnConnStart钩子函数的方法
	CallOnConnStart(conneciton IConneciton)
	//调用OnConnStop钩子函数的方法
	CallOnConnStop(conneciton IConneciton)
}

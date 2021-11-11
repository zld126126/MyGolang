package main


import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	h      map[string]http.Handler
	addr   string
	server *http.Server
}

func InitHandler(addr string) *Handler {
	return &Handler{
		h:    make(map[string]http.Handler),
		addr: addr,
	}
}

func (p *Handler) Add(url string, handler http.Handler) {
	p.h[url] = handler
}

func (p *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	urlPath := req.URL.Path
	if h, ok := p.h[urlPath]; ok {
		fmt.Println(req.Method," ===> ", p.addr+urlPath)
		h.ServeHTTP(resp, req)
		return
	}

	// NOTFOUND 处理
	http.NotFound(resp, req)
}

func (p *Handler) Mount(m map[string]http.Handler) {
	for urlPath, handler := range m {
		p.Add(urlPath, handler)
	}
}

func (p *Handler) GetServer() *http.Server {
	return p.server
}

func (p *Handler) showApi() {
	for urlPath, _ := range p.h {
		fmt.Println(p.addr + urlPath)
	}
}

func (p *Handler) Run() {
	p.server = &http.Server{
		Addr:    p.addr,
		Handler: p,
	}

	p.showApi()

	fmt.Println("server is run:", p.addr)
	if err := p.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

/**
main.go 调用示例:
	package main

	import (
		"fmt"
		"net/http"
	)

	type Controller struct{}

	func (p *Controller) ServeHTTP(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("server"))
	}

	// 自封装-http请求
	func main() {
		addr := "127.0.0.1" + ":" + fmt.Sprint(8080)
		handler := InitHandler(addr)
		handler.Mount(map[string]http.Handler{
			"/":    &Controller{},
			"/api": &Controller{},
		})

		// 测试api 浏览器|curl:
		// curl http://127.0.0.1:8080/
		// curl http://127.0.0.1:8080/api
		// 启动:
		handler.Run()
	}
*/

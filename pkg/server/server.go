package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/fyuan1316/proxyserver/cmd/option"
)

const (
	swaggerFilesPath    = "dist"
	swaggerJSONTemplate = "dist/swagger-template.json"
	swaggerJSONFile     = "./dist/swagger.json"
)

// MyServer MyServer
type MyServer struct {
	Configuration *option.Configuration
	*http.Server
}

// NewMyServer createServer
func NewMyServer(addr string, options ...func(*option.Configuration) error) *MyServer {
	server := &MyServer{
		Configuration: &option.Configuration{},
		Server:        &http.Server{},
	}
	server.Configuration.SetAddr(addr)
	for _, opt := range options {
		opt(server.Configuration)
	}
	return server
}

//Start Start Server
func (s *MyServer) Start() {
	go func() {
		s.preprocess(swaggerJSONTemplate, swaggerJSONFile)
		fmt.Printf("swagger server start listen on %v\n", s.Configuration.Addr)
		s.Server.Handler = s.getHandlers()
		s.Server.Addr = s.Configuration.Addr
		if err := s.Server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	fmt.Println("swagger server shutdown now")
	s.Server.Shutdown(context.Background())

}

func (s *MyServer) hanlder(rw http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(s.Configuration.TargetURL) // already checked
	proxy := httputil.NewSingleHostReverseProxy(url)
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	fmt.Println(req.URL.Host)
	fmt.Println(req.URL.Scheme)
	fmt.Println(req.Host)
	fmt.Println(req.URL.Port)
	proxy.ServeHTTP(rw, req)
}

func (s *MyServer) getHandlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(s.Configuration.ProxyPrefixURLCondition, s.hanlder)
	mux.Handle("/", http.FileServer(http.Dir(swaggerFilesPath)))
	return mux
}

func (s *MyServer) preprocess(fromTemplatePath, toSwaggerfile string) {
	RenderTemplate(s.Configuration.ModelName, fromTemplatePath, toSwaggerfile)
}

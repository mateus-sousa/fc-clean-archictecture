package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string][]MethodHandlerFunc
	WebServerPort string
}

type MethodHandlerFunc struct {
	Method string
	Handle http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string][]MethodHandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc, method string) {
	s.Handlers[path] = append(s.Handlers[path], MethodHandlerFunc{
		Method: method,
		Handle: handler,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		for _, h := range handler {
			s.Router.Method(h.Method, path, h.Handle)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}

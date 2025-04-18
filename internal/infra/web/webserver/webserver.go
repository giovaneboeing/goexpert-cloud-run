package webserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (w *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	w.Handlers[path] = handler
}

func (w *WebServer) Start() {
	w.Router.Use(middleware.Logger)

	for path, handler := range w.Handlers {
		w.Router.Handle(path, handler)
	}

	fmt.Println("Iniciando servidor na porta", w.WebServerPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", w.WebServerPort), w.Router); err != nil {
		panic(err)
	}
}

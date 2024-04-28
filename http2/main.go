package main

import (
	"fmt"
	"net/http"
)

type CustomServer struct {
	// server
	getPath2Handler  map[string][]HTTPHandler
	postPath2Handler map[string][]HTTPHandler
}

type ServerHTTPHanlder struct {
	server CustomServer
}

type HTTPHandler func(w http.ResponseWriter, req *http.Request)

func (server *CustomServer) Get(pathname string, handler HTTPHandler) {
	if server.getPath2Handler == nil {
		server.getPath2Handler = make(map[string][]HTTPHandler)
	}

	var exist = server.getPath2Handler[pathname]

	if exist != nil {
		server.getPath2Handler[pathname] = append(server.getPath2Handler[pathname], handler)
	} else {
		server.getPath2Handler[pathname] = []HTTPHandler{handler}
	}

}

func (server *CustomServer) Post(pathname string, handler HTTPHandler) {
	if server.postPath2Handler == nil {
		server.postPath2Handler = make(map[string][]HTTPHandler)
	}

	var exist = server.postPath2Handler[pathname]

	if exist != nil {
		server.postPath2Handler[pathname] = append(server.postPath2Handler[pathname], handler)
	} else {
		server.postPath2Handler[pathname] = []HTTPHandler{}
	}
}

func (m *ServerHTTPHanlder) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var path = req.URL.Path
	var method = req.Method

	fmt.Printf("%s", m.server.getPath2Handler)

	if method == "GET" {
		var handlers = m.server.getPath2Handler[path]

		if handlers != nil {
			for _, handler := range handlers {
				handler(w, req)
			}
		}
	} else if method == "POST" {
		var handlers = m.server.postPath2Handler[path]

		if handlers != nil {
			for _, handler := range handlers {
				handler(w, req)
			}
		}
	}
}

func (server *CustomServer) New(port string) {
	handler := &ServerHTTPHanlder{server: *server}

	fmt.Println("Server is running on port " + port)

	http.ListenAndServe(":"+port, handler)
}

func main() {
	server := &CustomServer{}

	server.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	server.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello"))
	})

	server.Get("/world", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	server.New("8080")
}

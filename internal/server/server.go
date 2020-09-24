package server

import (
	"context"
	"log"
	"net/http"
	"rt/api"
	http2 "rt/internal/rt/delivery/http"
	"rt/internal/rt/repository"
	"rt/internal/rt/usecase"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Server describes file server struct
type Server struct {
	router      *mux.Router
	server      *http.Server
	shutdownReq chan bool
}

// New gets Apps params
func New() *Server {

	return &Server{
		router:      mux.NewRouter(),
		shutdownReq: make(chan bool),
	}
}

// Start method start test service
func (s *Server) start(host string, port int) {

	api.MakeProductHandlers(s.router, s.shutdownReq, http2.RT{UC: usecase.NewRtUseCase(repository.NewRtRepository())})

	s.server = &http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: s.router,
	}

	go func() {
		log.Println("Server startig...")
		err := s.server.ListenAndServe()
		if err != nil {
			log.Println("Listen and serve:", err)
		}

	}()
	s.waitShutdown()
}

//waitShutdown waits shutdownReq
func (s *Server) waitShutdown() {

	<-s.shutdownReq
	log.Println("Shutdown...")

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)

	if err != nil {
		log.Println("Shutdown request error:", err)
	}

	log.Println("Stopping http file server...")

}

// Start runs http-server with the specified port
func Start(host string, port int) {

	srv := New()
	srv.start(host, port)
}

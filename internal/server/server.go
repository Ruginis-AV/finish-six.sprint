package server

import (
	"log"
	"net/http"
	"time"
)

type handler interface {
	HtmlHandler(w http.ResponseWriter, r *http.Request)
	ParseHandler(w http.ResponseWriter, r *http.Request)
}
type Server struct {
	server *http.Server
	logger *log.Logger
}

func New(
	logger *log.Logger,
	addr string,
	port string,
	writeTimeOut int,
	readTimeOut int,
	idleTimeOut int,
	handler handler,
) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HtmlHandler)
	mux.HandleFunc("/upload", handler.ParseHandler)

	server := &http.Server{
		Handler:      mux,
		Addr:         addr + ":" + port,
		WriteTimeout: time.Duration(writeTimeOut) * time.Second,
		ReadTimeout:  time.Duration(readTimeOut) * time.Second,
		IdleTimeout:  time.Duration(idleTimeOut) * time.Second,
		ErrorLog:     logger,
	}

	return &Server{
		server: server,
		logger: logger,
	}
}

func (s *Server) Start() error {
	s.logger.Printf("Server Start on %v\n", s.server.Addr)
	return s.server.ListenAndServe()
}

package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
	Host   string
	Port   string
}

// NewServer возвращает экземпляр структуры сервера
func NewServer() *Server {
	return &Server{}
}

// Run метод конфигурирует сервер и возвращает результат выполнения ListenAndServe. В случае ошибки запуска сервера
// она будет обработана извне при вызове метода.
func (s *Server) Run(host, port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr:           host + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}
	return s.server.ListenAndServe()
}

// Shutdown позволяет реализовать graceful shutdown, для "мягкой" остановки работы сервера
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

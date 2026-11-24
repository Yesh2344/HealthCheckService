package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"healthcheckservice/config"
	"healthcheckservice/health"
	"healthcheckservice/logger"
	"healthcheckservice/monitor"
)

// Server encapsulates the HTTP server and its dependencies.
type Server struct {
	cfg     *config.Config
	monitor *monitor.Monitor
	router  *mux.Router
}

// New creates and configures a new Server instance.
func New(cfg *config.Config, mon *monitor.Monitor) *Server {
	s := &Server{
		cfg:     cfg,
		monitor: mon,
		router:  mux.NewRouter(),
	}
	s.routes()
	return s
}

// routes registers all HTTP routes.
func (s *Server) routes() {
	s.router.HandleFunc("/health", health.Handler(s.monitor)).Methods(http.MethodGet)
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Health Check Monitoring Service")
	})
}

// Start launches the HTTP server.
func (s *Server) Start() error {
	addr := ":" + s.cfg.Port
	logger.Info("Starting server on ", addr)
	return http.ListenAndServe(addr, s.router)
}
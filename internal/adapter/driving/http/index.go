package http

import (
	"net/http"
	"time"

	"github.com/FinalProjectTMS/authservice/internal/config"
	"github.com/FinalProjectTMS/authservice/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	uc     *usecase.UseCases
}

const httpServerReadHeaderTimeout = 70 * time.Second

func New(
	cfg *config.Config,
	uc *usecase.UseCases,
) *http.Server {
	r := gin.New()

	// FOR DEVELOPMENT PURPOSES ONLY
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization", "X-Refresh-Token"},
	}))

	srv := &Server{
		router: r,
		cfg:    cfg,
		uc:     uc,
	}

	srv.endpoints()

	httpServer := &http.Server{
		Addr:              cfg.HTTPPort,
		Handler:           srv,
		ReadHeaderTimeout: httpServerReadHeaderTimeout,
	}

	// srv.log.Info(fmt.Sprintf("HTTP server is initialized on port: %v", cfg.HTTPPort))

	return httpServer
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

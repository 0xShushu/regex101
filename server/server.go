package server

import (
	"github.com/0xshushu/regex101/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

//function to initialize server
func NewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

//function to mount handlers
func (s *Server) MountHandlers() {
	//set middleware
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.RedirectSlashes)

	//set routes
	s.Router.Get("/", handlers.Index)
	s.Router.NotFound(handlers.NotFound)
}
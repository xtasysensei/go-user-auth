package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/xtasysensei/go-poll/pkg/handlers"
	"github.com/xtasysensei/go-poll/pkg/handlers/user"
)

func RegisterRoutes(apiRouter *chi.Mux) {

	apiRouter.Get("/", handlers.Index)
	apiRouter.Get("/ping", handlers.Health)

	apiRouter.Route("/v1", func(route chi.Router) {
		//user routes
		route.Route("/auth", func(r chi.Router) {
			r.Post("/login", user.HandleLogin)
			r.Post("/register", user.HandleRegister)
		})

	})

}

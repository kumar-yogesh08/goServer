package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	middlewar "goServer/internal/middleware"
)

func Handler(r *chi.Mux) {

	r.Use(middleware.StripSlashes)

	r.Route("/account",func(router chi.Router) {
		router.Use(middlewar.Authorization)
		router.Get("/coins",GetCoinBalance)
	})

}

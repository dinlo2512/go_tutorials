package handlers
import (
   "github.com/go-chi/chi"
   chimiddle  "github.com/go-chi/chi/middleware"
    "GOAPI/internal/middleware"
)

func Handler(r *chi.Mux) {
    r.Use(chimiddle.StripSlashes)

    r.Route("/account", func(route chi.Router) {
       router.Use(middleware.Authorization)

       route.Get("/coins", GetCoinBalance)

    })

}
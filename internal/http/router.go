package http

import (
	"net/http"
	"starter-golang/internal/handler"
	intmw "starter-golang/internal/http/middleware"

	chi "github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)


func NewRouter() http.Handler {
r := chi.NewRouter()
r.Use(chimw.RequestID)
r.Use(chimw.RealIP)
r.Use(intmw.Recoverer)
r.Use(intmw.RequestLogger)
r.Use(intmw.CORS())


r.Get("/health", handler.Health)


// Kelompokkan versi
r.Route("/v1", func(r chi.Router) {
// contoh: r.Mount("/users", usersRouter())
})


return r
}
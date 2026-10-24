package application

import (
	"net/http"

	"github.com/abhinaaaavvv/ilovego/OrderAPI/handler"
	"github.com/abhinaaaavvv/ilovego/OrderAPI/repository/order"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *App) loadRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	orderHandler := &handler.Order{
		Repo: &order.RedisRepo{
			Client: app.rdb,
		},
	}

	r.Route("/orders", func(r chi.Router) {
		r.Post("/", orderHandler.Create)
		r.Get("/", orderHandler.List)
		r.Get("/{id}", orderHandler.GetByID)
		r.Put("/{id}", orderHandler.UpdateByID)
		r.Delete("/{id}", orderHandler.DeleteByID)
	})

	app.router = r
}

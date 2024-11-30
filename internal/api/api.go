package api

import (
	"net/http"

	"github.com/brnocorreia/api-golang-template/internal/sample/samplehttp"
	"github.com/brnocorreia/api-golang-template/internal/sample/sampleservice"
	"github.com/brnocorreia/api-golang-template/internal/store/pgstore"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

const (
	ApiVersion = "api/v1"
)

type ApiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := ApiHandler{
		q: q,
	}
	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		// TODO: make this configurable
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	sampleService := sampleservice.New(q)
	sampleHttp := samplehttp.New(sampleService)
	samplehttp.InitSampleRoutes(r, sampleHttp)

	a.r = r
	return a
}
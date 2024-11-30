package samplehttp

import (
	"github.com/brnocorreia/api-golang-template/internal/config"
	"github.com/go-chi/chi/v5"
)

func InitSampleRoutes(r *chi.Mux, h *sampleHttp) {
	r.Route(config.ApiVersion, func(r chi.Router) {
		r.Route("/sample", func(r chi.Router) {
			r.Post("/", h.CreateSampleHandler)
			
			r.Route("/{sampleId}", func(r chi.Router) {
				r.Get("/", h.GetSampleByIDHandler)
			})
		})
	})
}
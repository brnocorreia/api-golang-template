package samplehttp

import (
	"net/http"

	"github.com/brnocorreia/api-golang-template/internal/sample/samplecoder"
	"github.com/brnocorreia/api-golang-template/internal/sample/sampleservice"
	"github.com/brnocorreia/api-golang-template/internal/utils"
)

type SampleHttp interface {}

type sampleHttp struct {
	sampleService sampleservice.SampleService
}

func New(sampleService sampleservice.SampleService) *sampleHttp {
	return &sampleHttp{
		sampleService: sampleService,
	}
}

func (h *sampleHttp) GetSampleByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := samplecoder.DecodeSampleIDFromURI(r)
	if err != nil {
		utils.SendJSON(w, nil, false, "Unable to decode sample id", http.StatusBadRequest)
		return
	}
	sample, err := h.sampleService.GetByID(ctx, id.String())
	if err != nil {
		utils.SendJSON(w, nil, false, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJSON(w, sample, true, "OK", http.StatusOK)
}

func (h *sampleHttp) CreateSampleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sample, err := samplecoder.DecodeSampleFromBody(r)
	if err != nil {
		utils.SendJSON(w, nil, false, "Unable to decode sample", http.StatusBadRequest)
		return
	}
	newSample, err := h.sampleService.Create(ctx, sample)
	if err != nil {
		utils.SendJSON(w, nil, false, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJSON(w, newSample, true, "OK", http.StatusOK)
}
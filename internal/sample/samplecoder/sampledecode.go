package samplecoder

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/brnocorreia/api-golang-template/internal/sample/sampledomain/sampleentities"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)


func DecodeSampleIDFromURI(r *http.Request) (uuid.UUID, error) {
	id := chi.URLParam(r, "sampleId")
	sampleUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, errors.New("invalid sample id")
	}
	return sampleUUID, nil
}

func DecodeSampleFromBody(r *http.Request) (*sampleentities.CreateSamplePayload, error) {
	createSample := &sampleentities.CreateSamplePayload{}
	err := json.NewDecoder(r.Body).Decode(&createSample)
	if err != nil {
		return nil, err
	}
	return createSample, nil
}
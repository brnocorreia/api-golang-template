package sampleservice

import (
	"context"
	"errors"
	"time"

	"github.com/brnocorreia/api-golang-template/internal/sample/samplecoder"
	"github.com/brnocorreia/api-golang-template/internal/sample/sampledomain/sampleentities"
	"github.com/brnocorreia/api-golang-template/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type SampleService interface {
	GetByID(ctx context.Context, id string) (*sampleentities.Sample, error)
	Create(ctx context.Context, sample *sampleentities.CreateSamplePayload) (*sampleentities.Sample, error)
}

type sampleService struct {
	q *pgstore.Queries
}

func New(q *pgstore.Queries) *sampleService {
	return &sampleService{
		q: q,
	}
}

func (s *sampleService) GetByID(ctx context.Context, id string) (*sampleentities.Sample, error) {

	// ID must be a valid UUID
	sampleUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	sample, err := s.q.GetSampleByID(ctx, sampleUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("sample_not_found")
		}
		return nil, err
	}
	return samplecoder.EncodeSample(sample), nil
}

func (s *sampleService) Create(ctx context.Context, sample *sampleentities.CreateSamplePayload) (*sampleentities.Sample, error) {

	now := time.Now().UTC()

	newSample, err := s.q.CreateSample(ctx, pgstore.CreateSampleParams{
		Name:      sample.Name,
		Password:  sample.Password,
		CreatedAt: now,
	})
	if err != nil {
		return nil, err
	}

	return samplecoder.EncodeSample(newSample), nil
}
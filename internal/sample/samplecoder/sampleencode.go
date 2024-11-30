package samplecoder

import (
	"github.com/brnocorreia/api-golang-template/internal/sample/sampledomain/sampleentities"
	"github.com/brnocorreia/api-golang-template/internal/store/pgstore"
)


func EncodeSample(sample pgstore.Sample) (*sampleentities.Sample) {
	return &sampleentities.Sample{
		ID:        sample.ID.String(),
		Name:      sample.Name,
		Password:  sample.Password,
		CreatedAt: sample.CreatedAt,
	}
}
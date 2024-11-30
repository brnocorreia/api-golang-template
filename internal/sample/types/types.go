package types

type contextKey string

const (
	SampleID contextKey = "sampleId"
)

type CreateSamplePayload struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
}

type SamplePayload struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}


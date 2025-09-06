package reptiloid

type client[T any] struct {
	Model  *replicateModel[T]
	apikey string
}

// Generic model structure with input type T
type replicateModel[T any] struct {
	//Id of replicate model
	Id string
	//Price per 1 image or 1 million tokens
	Cost float64
}

type requestBody[T any] struct {
	Input T `json:"input"`
}

type responseOutput struct {
	Id          string   `json:"id"`
	Model       string   `json:"model"`
	Version     string   `json:"version"`
	Logs        string   `json:"logs"`
	Output      []string `json:"output"`
	DataRemoved bool     `json:"data_removed"`
	Error       *string  `json:"error"`
	Status      any      `json:"status"`
	CreatedAt   string   `json:"created_at"`
}

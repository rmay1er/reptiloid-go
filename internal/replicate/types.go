package replicate


type client[T any] struct {
	model  replicateModel[T]
	apikey string
}

// Generic model structure with input type T
type replicateModel[T any] struct {
	id string
}

type requestBody[T any] struct {
	Input T `json:"input"`
}

type responseOutput struct {
 Id        string   `json:"id"`
 Model     string   `json:"model"`
 Version   string   `json:"version"`
 Logs      string   `json:"logs"`
 Output    []string `json:"output"`
 DataRemoved bool   `json:"data_removed"`
 Error     *string  `json:"error"`
 Status    int   `json:"status"`
 CreatedAt string   `json:"created_at"`
}

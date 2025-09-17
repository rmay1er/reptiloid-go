package reptiloid

import (
	"bytes"
	"encoding/json/v2"
	"io"
	"net/http"

	"github.com/rmay1er/reptiloid-go/internal/utils"
)

// NewClient creates a new client with the specified replicate model and API key.
func NewClient[T any](model *ReplicateModel[T], apikey string) *Client[T] {
	return &Client[T]{Model: model, apikey: apikey}
}

// NewReplicateModel creates a new replicateModel instance with the given model ID.
// Optionally, a cost value can be provided; if not, the cost defaults to zero.
func NewReplicateModel[T any](id string, cost ...float64) ReplicateModel[T] {
	model := ReplicateModel[T]{Id: id}
	if len(cost) > 0 && cost[0] != 0 {
		model.Cost = cost[0]
	}
	return model
}

// Generate sends a request to generate an image based on the input parameters.
// It constructs the HTTP POST request, adds necessary headers, and parses the response.
func (c *Client[T]) Generate(input *T) (responseOutput, error) {

	// Получаем ID модели
	id := c.Model.Id

	// Создаём ссылку на модель
	url, err := utils.CreateApiReqUrl(id)
	if err != nil {
		return responseOutput{}, err
	}

	// Формируем тело запроса
	body := requestBody[T]{
		Input: *input,
	}

	// Преобразуем тело в JSON
	jsonData, err := json.Marshal(body, json.OmitZeroStructFields(true))
	if err != nil {
		return responseOutput{}, err
	}
	// Создаем POST запрос с телом
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return responseOutput{}, err
	}

	// Добавляем заголовки к запросу
	req.Header.Set("Authorization", "Bearer "+c.apikey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "wait")

	// Создаем HTTP клиент
	client := &http.Client{}

	// Отправляем запрос через клиент
	resp, err := client.Do(req)
	if err != nil {
		return responseOutput{}, err
	}
	// Закрываем поток ответа после завершения функции
	defer resp.Body.Close()

	// Читаем тело ответа
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseOutput{}, err
	}

	// Парсим тело ответа в структуру
	var apiResp responseOutput
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		return responseOutput{}, err
	}

	return apiResp, nil
}

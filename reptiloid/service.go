package reptiloid

import (
	"bytes"
	"encoding/json/v2"
	"fmt"
	"github.com/rmay1er/reptiloid-go/internal/utils"
	"io"
	"net/http"
)

// NewClient creates a new client with the specified replicate model and API key.
func NewClient[T any](model replicateModel[T], apikey string) *client[T] {
	return &client[T]{model: model, apikey: apikey}
}

// NewReplicateModel creates a new replicate model with the given model ID.
func NewReplicateModel[T any](id string) replicateModel[T] {
	return replicateModel[T]{id: id}
}

// GenerateImage sends a request to generate an image based on the input parameters.
// It constructs the HTTP POST request, adds necessary headers, and parses the response.
func (c *client[T]) GenerateImage(input T) (responseOutput, error) {

	// Получаем ID модели
	id := c.model.id

	// Создаём ссылку на модель
	url, err := utils.CreateApiReqUrl(id)
	if err != nil {
		return responseOutput{}, err
	}

	// Формируем тело запроса
	body := requestBody[T]{
		Input: input,
	}

	// Преобразуем тело в JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		return responseOutput{}, err
	}

	fmt.Println(string(jsonData))

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

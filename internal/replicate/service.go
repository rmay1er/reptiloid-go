package replicate

import (
	"bytes"
	"encoding/json/v2"
	"fmt"
	"net/http"
	"io"
)

func NewClient[T any](model replicateModel[T], apikey string) *client[T] {
	return &client[T]{model: model, apikey: apikey}
}

func (c *client[T]) MockGenerateImage(input T) (string, error) {
	// Здесь логика обращения по API с использованием c.model и input
	result := fmt.Sprintf("Generating image for model %s with input: %+v\n", c.model, input)
	return result, nil
}

func createApiReqUrl(model string) (string, error) {
	modelUrl := fmt.Sprintf("https://api.replicate.com/v1/models/%s/predictions", model)
	return modelUrl, nil
}

func NewReplicateModel[T any](id string) replicateModel[T] {
	return replicateModel[T]{id: id}
}

func (c *client[T]) GenerateImage(input T) (responseOutput, error) {

	// Получаем ID модели
	id := c.model.id

	// Создаём ссылку на модель
	url, err := createApiReqUrl(id)
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

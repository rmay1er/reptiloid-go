package utils

import (
	"fmt"
)

func CreateApiReqUrl(model string) (string, error) {
	modelUrl := fmt.Sprintf("https://api.replicate.com/v1/models/%s/predictions", model)
	return modelUrl, nil
}

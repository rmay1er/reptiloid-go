package replicate

import (
	"testing"
)

func TestCreateApiReqUrl(t *testing.T) {
	model := "black-forest-labs/flux-schnell"
	got, err := createApiReqUrl(model)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	t.Logf("URL для модели %q: %s", model, got)
}

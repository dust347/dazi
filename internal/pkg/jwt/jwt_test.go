package jwt

import (
	"testing"

	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/config"
)

func TestSign(t *testing.T) {
	config.SetConfig4Test(&model.Config{
		JWT: model.JWTConfig{
			SignKey: "test",
		},
	})
	token, err := Sign("12345")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)

	id, valid := Parse(token)
	t.Log(id, valid)
}

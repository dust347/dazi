package wx

import (
	"context"
	"testing"

	"github.com/dust347/dazi/internal/model"
)

func TestCheck(t *testing.T) {
	cli, err := NewLoginCheckClient(&model.DatabaseConfig{})

	if err != nil {
		t.Fatal(err)
	}

	resp, err := cli.Check(context.Background(), &model.LoginCheckReq{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("resp: %+v", resp)
}

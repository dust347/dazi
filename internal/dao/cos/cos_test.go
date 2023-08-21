package cos

import (
	"bytes"
	"context"
	"testing"

	"github.com/dust347/dazi/internal/model"
)

func TestUpload(t *testing.T) {
	cli, err := NewClient(&model.DatabaseConfig{})
	if err != nil {
		t.Fatal(err)
	}

	r := bytes.NewReader([]byte("test"))
	path, err := cli.Upload(context.Background(), "test/test.txt", r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(path)
}

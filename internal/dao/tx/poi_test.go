package tx

import (
	"context"
	"testing"

	"github.com/dust347/dazi/internal/model"
)

func TestGetCity(t *testing.T) {
	cli, err := NewPoiClient(&model.DatabaseConfig{
		Namespace: "5QTBZ-OQK6U-KELVF-4FXQ4-Q46T3-WWFWZ",
		Target:    "K7weLWA0bHgqnThjQhET7fIoGGeQ6Hn8",
	})
	if err != nil {
		t.Fatal(err)
	}

	city, err := cli.GetCity(context.Background(), &model.Location{
		Latitude:  39915003,
		Longitude: 116483574,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("city: %+v", city)
}

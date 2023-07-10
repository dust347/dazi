package config

import "testing"

func TestLoad(t *testing.T) {
	// log
	if err := Load(); err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", GetConfig())
}

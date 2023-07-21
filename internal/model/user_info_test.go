package model

import (
	"encoding/json"
	"testing"
	"time"
)

func TestUserInfoJSON(t *testing.T) {
	u := UserInfo{
		OpenID:   "123",
		Phone:    "18511112222",
		Birthday: Date(time.Now()),
		Gender:   GenderMale,
		City:     "1",
		CityName: "北京",
		NickName: "nick",
		Tags:     Tags{"吃", "摸鱼"},
		Location: Location{
			Latitude:  10000,
			Longitude: 20000,
		},
	}

	b, err := json.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)

	var user UserInfo
	if err := json.Unmarshal(b, &user); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", user)
}

func TestLocation(t *testing.T) {
	loc := Location{
		Latitude:  39915003,
		Longitude: 116483574,
	}
	t.Log(loc.String())

	b, err := json.Marshal(loc)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
}

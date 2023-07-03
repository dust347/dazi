package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/dust347/dazi/internal/model"
)

func TestCreate(t *testing.T) {
	cli, err := NewUserInfoClient(&model.DatabaseConfig{
		Type:   model.DatabaseTypeMysql,
		Target: "root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		Name:   "t_user_info",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cli.Create(context.Background(), &model.UserInfo{
		ID:       "12345",
		OpenID:   "234",
		Phone:    "18522223333",
		Birthday: model.Date(time.Now()),
		Gender:   model.GenderMale,
		City:     "156110000",
		CityName: "北京市",
		NickName: "刘馋",
		Tags:     model.Tags{"喵", "吃"},
		Location: model.Location{
			Longitude: 39915003,
			Latitude:  116483574,
		},
	}))
}

package mysql

import (
	"context"
	"net"
	"testing"

	"github.com/dust347/dazi/internal/model"
)

func TestHost(t *testing.T) {
	hp := net.JoinHostPort("bj-cdb-gfmnayaq.sql.tencentcdb.com", "63814")
	t.Log(hp)
}

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
		ID:     "12345678",
		OpenID: "234568",
		Phone:  "18522223333",
		//Birthday: model.Date(time.Now()),
		Gender:   model.GenderMale,
		City:     "156110000",
		CityName: "北京市",
		NickName: "刘肉段",
		Tags:     model.Tags{"喵", "吃"},
		Location: model.Location{
			Longitude: 39915113,
			Latitude:  116484554,
		},
	}))
}

func TestQuery(t *testing.T) {
	cli, err := NewUserInfoClient(&model.DatabaseConfig{
		Type:   model.DatabaseTypeMysql,
		Target: "root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		Name:   "t_user_info",
	})
	if err != nil {
		t.Fatal(err)
	}

	user, err := cli.Query(context.Background(), "123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("user: %+v", user)
}

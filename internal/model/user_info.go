package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// UserInfo 用户信息
type UserInfo struct {
	ID          string `json:"id" gorm:"column:id"`
	IdentityNum string `json:"identity_number" gorm:"column:identity_number"`
	Phone       string `json:"phone" gorm:"column:phone"`

	Birthday time.Time `json:"birthday" gorm:"column:birthday"`
	Gender   Gender    `gorm:"column:gender"`
	City     string    `gorm:"column:city"`
	NickName string    `gorm:"column:nick_name"`
	Tags     Tags      `gorm:"column:tags"`
	Location Location  `gorm:"column:location"`
}

// Tags 用户标签
type Tags []string

// Scan 实现 sql.Scanner 接口
func (tags *Tags) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to unmarshal, value: %+v", value)
	}

	*tags = strings.Split(s, ",")
	return nil
}

// Value 实现 driver.Valuer 接口
func (tags Tags) Value() (driver.Value, error) {
	return strings.Join(tags, ","), nil
}

// Gender 性别
type Gender int8

const (
	// GenderUnknown 未知
	GenderUnknown = 0
	// GenderMale 男
	GenderMale = 1
	// GenderFemale 女
	GenderFemale = 2
)

// Location 用户位置
type Location struct {
	Latitude  int64
	Longitude int64
}

// Scan 实现 sql.Scanner 接口
func (loc *Location) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to unmarshal, value: %+v", value)
	}

	l := strings.Split(s, ",")
	if len(l) != 2 {
		return fmt.Errorf("failed to unmarshal, value: %+v", value)
	}

	lat, err := strconv.ParseInt(l[0], 10, 64)
	if err != nil {
		return err
	}
	lon, err := strconv.ParseInt(l[1], 10, 64)
	if err != nil {
		return err
	}

	loc.Latitude = lat
	loc.Longitude = lon
	return nil
}

// Value 实现 driver.Valuer 接口
func (loc Location) Value() (driver.Value, error) {
	return fmt.Sprintf("%d,%d", loc.Latitude, loc.Longitude), nil
}

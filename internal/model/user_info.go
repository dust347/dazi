package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     string `json:"id" gorm:"column:id"`
	OpenID string `json:"open_id" gorm:"column:open_id"`
	Phone  string `json:"phone" gorm:"column:phone"`

	Birthday  Date     `json:"birthday" gorm:"column:birthday"`
	Gender    Gender   `json:"gender" gorm:"column:gender"`
	City      string   `json:"city" gorm:"column:city"`
	CityName  string   `json:"city_name" gorm:"column:city_name"`
	NickName  string   `json:"nick_name" gorm:"column:nick_name"`
	AvatarURL string   `json:"avatar_url" gorm:"column:avatar_url"`
	Tags      Tags     `json:"tags" gorm:"column:tags"`
	Location  Location `json:"location" gorm:"column:location"`
	Detail    string   `json:"detail" gorm:"column:detail"`
}

// Date 日期
type Date time.Time

// DateFormat ...
const DateFormat = "2006-01-02"

// UnmarshalJSON ...
func (date *Date) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+DateFormat+`"`, string(b), time.Local)
	if err != nil {
		return err
	}
	*date = Date(now)
	return nil
}

// MarshalJSON ...
func (date Date) MarshalJSON() ([]byte, error) {
	if time.Time(date).IsZero() {
		return []byte("\"\""), nil
	}
	b := make([]byte, 0, len(DateFormat)+2)
	b = append(b, '"')
	b = time.Time(date).AppendFormat(b, DateFormat)
	b = append(b, '"')
	return b, nil
}

// String ...
func (date Date) String() string {
	if time.Time(date).IsZero() {
		return ""
	}
	return time.Time(date).Format(DateFormat)
}

// Value insert timestamp into mysql need this function.
func (date Date) Value() (driver.Value, error) {
	if time.Time(date).IsZero() {
		return nil, nil
	}
	return time.Time(date), nil
}

// Scan value of time.Time
func (date *Date) Scan(value interface{}) error {
	timeValue, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal time value:", value))
	}

	if timeValue.IsZero() {
		*date = Date(time.Time{})
		return nil
	}

	*date = Date(timeValue)
	return nil
}

// Tags 用户标签
type Tags []string

// Scan 实现 sql.Scanner 接口
func (tags *Tags) Scan(value interface{}) error {
	s, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal, value: %+v", value)
	}

	*tags = strings.Split(string(s), ",")
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
	Latitude  int64 `json:"latitude"`
	Longitude int64 `json:"longitude"`
}

// IsEmpty 是否为空
func (loc *Location) IsEmpty() bool {
	return loc.Latitude == 0 && loc.Longitude == 0
}

// String ...
func (loc *Location) String() string {
	return fmt.Sprintf("%f,%f", float64(loc.Latitude)/1e6, float64(loc.Longitude)/1e6)
}

// Scan 实现 sql.Scanner 接口
func (loc *Location) Scan(value interface{}) error {
	s, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal, value: %+v", value)
	}

	l := strings.Split(string(s), ",")
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

// MarshalJSON ...
func (loc Location) MarshalJSON() ([]byte, error) {

	l := struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{
		Latitude:  float64(loc.Latitude) / 1e6,
		Longitude: float64(loc.Longitude) / 1e6,
	}
	return json.Marshal(l)
}

// UnmarshalJSON ...
func (loc *Location) UnmarshalJSON(data []byte) error {
	l := struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{}

	if err := json.Unmarshal(data, &l); err != nil {
		return err
	}
	loc.Latitude = int64(l.Latitude * 1e6)
	loc.Longitude = int64(l.Longitude * 1e6)
	return nil
}

// CityInfo 城市信息
type CityInfo struct {
	CityCode string
	CityName string
}

// ImageFile 图片文件
type ImageFile interface {
	io.Reader
}

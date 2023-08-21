package dao

import (
	"context"

	"github.com/dust347/dazi/internal/dao/cos"
	"github.com/dust347/dazi/internal/dao/mysql"
	"github.com/dust347/dazi/internal/dao/tx"
	"github.com/dust347/dazi/internal/dao/wx"
	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/errors"
)

// UserInfoManager 用户信息管理
type UserInfoManager interface {
	// 创建用户
	Create(ctx context.Context, user *model.UserInfo) error
	// 更新用户信息
	Update(ctx context.Context, user *model.UserInfo) error
	// 查询用户信息
	Query(ctx context.Context, id string) (*model.UserInfo, error)
	// 根据城市查找用户
	QueryUsersByCity(ctx context.Context, city string) ([]model.UserInfo, error)
	// 根据 openid 查询用户
	QueryByOpenID(ctx context.Context, openid string) (*model.UserInfo, error)
}

// NewUserInfoManager 创建 UserInfoManager 实例
func NewUserInfoManager(cfg *model.DatabaseConfig) (UserInfoManager, error) {
	if cfg == nil {
		return nil, errors.New(errors.ParamErr, "cfg is nil")
	}

	switch cfg.Type {
	case model.DatabaseTypeMysql:
		return mysql.NewUserInfoClient(cfg)
	default:
		return nil, errors.Errorf(errors.ParamErr, "not support type: %s", cfg.Type)
	}
}

// LoginChecker 登录校验接口
type LoginChecker interface {
	Check(ctx context.Context, req *model.LoginCheckReq) (*model.LoginCheckResp, error)
}

// NewLoginChecker 创建 LoginChecker 实例
func NewLoginChecker(cfg *model.DatabaseConfig) (LoginChecker, error) {
	if cfg == nil {
		return nil, errors.New(errors.ParamErr, "cfg is nil")
	}

	switch cfg.Type {
	case model.DatabaseTypeWxMiniProgram:
		return wx.NewLoginCheckClient(cfg)
	default:
		return nil, errors.Errorf(errors.ParamErr, "not support type: %s", cfg.Type)
	}
}

// PoiGetter poi 获取
type PoiGetter interface {
	GetCity(ctx context.Context, loc *model.Location) (*model.CityInfo, error)
}

// NewPoiGetter 创建 PoiGetter 实例
func NewPoiGetter(cfg *model.DatabaseConfig) (PoiGetter, error) {
	if cfg == nil {
		return nil, errors.New(errors.ParamErr, "cfg is nil")
	}

	switch cfg.Type {
	case model.DatabaseTypeTxMap:
		return tx.NewPoiClient(cfg)
	default:
		return nil, errors.Errorf(errors.ParamErr, "not support type: %s", cfg.Type)
	}
}

// ImageUploader 图片上报
type ImageUploader interface {
	Upload(ctx context.Context, fileName string, image model.ImageFile) (string, error)
}

// NewImageUploader 创建 ImageUploader 实例
func NewImageUploader(cfg *model.DatabaseConfig) (ImageUploader, error) {
	if cfg == nil {
		return nil, errors.New(errors.ParamErr, "cfg is nil")
	}

	switch cfg.Type {
	case model.DatabaseTypeCos:
		return cos.NewClient(cfg)
	default:
		return nil, errors.Errorf(errors.ParamErr, "not support type: %s", cfg.Type)
	}
}

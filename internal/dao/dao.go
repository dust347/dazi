package dao

import (
	"context"

	"github.com/dust347/dazi/internal/dao/mysql"
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

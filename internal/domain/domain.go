package domain

import (
	"context"

	"github.com/dust347/dazi/internal/model"
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
func NewUserInfoManager() (UserInfoManager, error) {
	return NewUserInfoRepo()
}

package domain

import (
	"context"

	"github.com/dust347/dazi/internal/model"
)

// UserInfoManager 用户信息管理
type UserInfoManager interface {
	Login(ctx context.Context, jsCode string, user *model.UserInfo) (*model.UserInfo, error)
	// 创建用户
	Create(ctx context.Context, user *model.UserInfo) error
	// 更新用户信息
	Update(ctx context.Context, user *model.UserInfo) error
	// 查询用户信息
	Query(ctx context.Context, id string) (*model.UserInfo, error)
	// 附近用户
	Nearby(ctx context.Context, id string, loc *model.Location) ([]model.UserInfo, error)
	// 更新头像
	UploadAvatar(ctx context.Context, userID, extName string, image model.ImageFile) (string, error)
}

// NewUserInfoManager 创建 UserInfoManager 实例
func NewUserInfoManager() (UserInfoManager, error) {
	return NewUserInfoRepo()
}

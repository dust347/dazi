package domain

import (
	"context"

	"github.com/dust347/dazi/internal/dao"
	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/config"
	"github.com/dust347/dazi/internal/pkg/errors"
	"github.com/dust347/dazi/internal/pkg/uuid"
)

// UserInfoRepo 用户信息管理
type UserInfoRepo struct {
	user dao.UserInfoManager
}

// NewUserInfoRepo 创建 UserInfoRepo 实例
func NewUserInfoRepo() (*UserInfoRepo, error) {
	var repo UserInfoRepo
	var err error

	repo.user, err = dao.NewUserInfoManager(&config.GetConfig().Database.UserInfo)
	if err != nil {
		return nil, errors.WithMsg(err, "init user info err")
	}

	return &repo, nil
}

// Create 创建用户信息
func (repo *UserInfoRepo) Create(ctx context.Context, user *model.UserInfo) error {
	if user == nil {
		return errors.New(errors.ParamErr, "user info is nil")
	}

	user.ID = uuid.New()
	if err := repo.user.Create(ctx, user); err != nil {
		return errors.WithMsg(err, "create user err")
	}

	return nil
}

// Update 更新用户信息
func (repo *UserInfoRepo) Update(ctx context.Context, user *model.UserInfo) error {
	if user == nil {
		return errors.New(errors.ParamErr, "")
	}

	if user.ID == "" {
		return errors.New(errors.ParamErr, "id empty")
	}

	if err := repo.user.Update(ctx, user); err != nil {
		return errors.WithMsg(err, "update err")
	}

	return nil
}

// Query 查询用户信息
func (repo *UserInfoRepo) Query(ctx context.Context, id string) (*model.UserInfo, error) {
	if id == "" {
		return nil, errors.New(errors.ParamErr, "id empty")
	}

	userInfo, err := repo.user.Query(ctx, id)
	if err != nil {
		return nil, errors.WithMsg(err, "query err")
	}

	return userInfo, nil
}

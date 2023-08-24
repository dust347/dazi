package mysql

import (
	"context"

	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserInfoClient 用户信息 client
type UserInfoClient struct {
	prx *gorm.DB
	cfg *model.DatabaseConfig
}

// NewUserInfoClient 创建 UserInfoClient 实例
func NewUserInfoClient(cfg *model.DatabaseConfig) (*UserInfoClient, error) {
	if cfg == nil {
		return nil, errors.New(errors.ParamErr, "cfg is nil")
	}

	db, err := gorm.Open(mysql.Open(cfg.Target), &gorm.Config{})
	if err != nil {
		return nil, errors.WithMsg(err, "open db err")
	}

	return &UserInfoClient{
		prx: db,
		cfg: cfg,
	}, nil
}

// Create 创建用户记录
func (cli *UserInfoClient) Create(ctx context.Context, user *model.UserInfo) error {
	if user == nil {
		return errors.New(errors.ParamErr, "user info is nil")
	}

	err := cli.prx.Table(cli.cfg.Name).Create(user).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return errors.Wrap(errors.DuplicatedErr, err, "duplicated user err")
		}
		return errors.WithMsg(err, "create err")
	}

	return nil
}

// Query 查询用户数据
func (cli *UserInfoClient) Query(ctx context.Context, id string) (*model.UserInfo, error) {
	var user model.UserInfo
	err := cli.prx.Table(cli.cfg.Name).
		Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.WithMsg(err, "select err")
	}

	return &user, nil
}

// QueryByOpenID 根据 openid 查询用户数据
func (cli *UserInfoClient) QueryByOpenID(ctx context.Context, openid string) (*model.UserInfo, error) {
	var user model.UserInfo
	err := cli.prx.Table(cli.cfg.Name).
		Where("open_id = ?", openid).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.WithMsg(err, "select err")
	}

	return &user, nil
}

// Update 更新数据
func (cli *UserInfoClient) Update(ctx context.Context, user *model.UserInfo) error {
	if user == nil {
		return errors.New(errors.ParamErr, "user is nil")
	}

	resp := cli.prx.Table(cli.cfg.Name).Where("id = ?", user.ID).Updates(user)
	if resp.Error != nil {
		return errors.WithMsg(resp.Error, "update err")
	}
	if resp.RowsAffected == 0 {
		return errors.New(errors.NoUserUpdateErr, "have no user update")
	}

	return nil
}

// QueryUsersByCity 根据城市码查询用户
func (cli *UserInfoClient) QueryUsersByCity(ctx context.Context, cityCode string) ([]model.UserInfo, error) {
	if cityCode == "" {
		return nil, nil
	}

	var users []model.UserInfo
	err := cli.prx.Table(cli.cfg.Name).
		Where("city = ?", cityCode).Find(&users).Error
	if err != nil {
		return nil, errors.WithMsg(err, "query users by city err")
	}

	return users, nil
}

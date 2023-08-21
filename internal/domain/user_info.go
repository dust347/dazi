package domain

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/dust347/dazi/internal/dao"
	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/config"
	"github.com/dust347/dazi/internal/pkg/errors"
	"github.com/dust347/dazi/internal/pkg/uuid"
)

// UserInfoRepo 用户信息管理
type UserInfoRepo struct {
	user  dao.UserInfoManager
	poi   dao.PoiGetter ``
	wx    dao.LoginChecker
	image dao.ImageUploader
}

// NewUserInfoRepo 创建 UserInfoRepo 实例
func NewUserInfoRepo() (*UserInfoRepo, error) {
	var repo UserInfoRepo
	var err error

	repo.user, err = dao.NewUserInfoManager(&config.GetConfig().Database.UserInfo)
	if err != nil {
		return nil, errors.WithMsg(err, "init user info err")
	}

	repo.poi, err = dao.NewPoiGetter(&config.GetConfig().Database.Poi)
	if err != nil {
		return nil, errors.WithMsg(err, "init poi err")
	}

	repo.wx, err = dao.NewLoginChecker(&config.GetConfig().Database.WxMiniProgram)
	if err != nil {
		return nil, errors.WithMsg(err, "init wx err")
	}

	repo.image, err = dao.NewImageUploader(&config.GetConfig().Database.Image)
	if err != nil {
		return nil, errors.WithMsg(err, "init image err")
	}

	return &repo, nil
}

// Login 登录逻辑
func (repo *UserInfoRepo) Login(ctx context.Context, jsCode string, user *model.UserInfo) (*model.UserInfo, error) {
	// wx 验证
	resp, err := repo.wx.Check(ctx, &model.LoginCheckReq{
		Code: jsCode,
	})
	if err != nil {
		return nil, errors.WithMsg(err, "wx login check err")
	}

	// 查询用户信息
	u, err := repo.user.QueryByOpenID(ctx, resp.OpenID)
	if err != nil {
		return nil, errors.WithMsg(err, "query user info err")
	}

	user.OpenID = resp.OpenID
	// 没有用户信息，则创建
	if u == nil {
		if err := repo.Create(ctx, user); err != nil {
			return nil, err
		}
		return user, nil
	}
	// 否则更新
	user.ID = u.ID
	if err := repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
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
	// 如果 location 不为空，获取地址信息
	if !user.Location.IsEmpty() {
		city, err := repo.poi.GetCity(ctx, &user.Location)
		if err != nil {
			return errors.WithMsg(err, "get city err")
		}
		user.City = city.CityCode
		user.CityName = city.CityName
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

// Nearby 附近的人
func (repo *UserInfoRepo) Nearby(ctx context.Context, id string, loc *model.Location) ([]model.UserInfo, error) {
	city, err := repo.poi.GetCity(ctx, loc)
	if err != nil {
		return nil, errors.WithMsg(err, "get city info err")
	}
	if city == nil {
		return nil, errors.New(errors.ParamErr, "city not found")
	}

	log.Printf("city: %+v", city)

	users, err := repo.user.QueryUsersByCity(ctx, city.CityCode)
	if err != nil {
		return nil, errors.WithMsg(err, "query users err")
	}
	if len(users) == 0 {
		return nil, nil
	}

	nearby := make([]model.UserInfo, 0, len(users))
	for i := range users {
		if users[i].ID == id {
			continue
		}
		nearby = append(nearby, users[i])
	}

	return nearby, nil
}

// UploadAvatar 上传头像
func (repo *UserInfoRepo) UploadAvatar(ctx context.Context, userID, extName string, image model.ImageFile) error {
	// 上传图片
	fileName := filepath.Join(userID, fmt.Sprintf("avatar%s", extName))
	path, err := repo.image.Upload(ctx, fileName, image)
	if err != nil {
		return errors.WithMsg(err, "upload image err")
	}

	// 更新图片地址
	err = repo.Update(ctx, &model.UserInfo{
		ID:        userID,
		AvatarURL: path,
	})
	if err != nil {
		return errors.WithMsg(err, "update avatar url err")
	}

	return nil
}

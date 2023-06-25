package service

import (
	"context"
	"net/http"

	"github.com/dust347/dazi/internal/domain"
	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/errors"
	"github.com/gin-gonic/gin"
)

var userInfo domain.UserInfoManager

// Init 初始化
func Init() error {
	var err error
	userInfo, err = domain.NewUserInfoManager()
	if err != nil {
		return errors.WithMsg(err, "init user info manager err")
	}

	return nil
}

// HTTPResp http response
type HTTPResp struct {
	ErrCode int32  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

// UserCreate 用户创建
func UserCreate(c *gin.Context) {
	var user model.UserInfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusOK, &HTTPResp{
			ErrCode: errors.ParamErr,
			ErrMsg:  "body unmarshal err",
		})
		return
	}

	if err := userInfo.Create(context.Background(), &user); err != nil {
		c.JSON(http.StatusOK, &HTTPResp{
			ErrCode: errors.Type(err),
			ErrMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &HTTPResp{})
	return
}

// UserUpdate 用户创建
func UserUpdate(c *gin.Context) {
	var user model.UserInfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusOK, &HTTPResp{
			ErrCode: errors.ParamErr,
			ErrMsg:  "body unmarshal err",
		})
		return
	}

	if err := userInfo.Update(context.Background(), &user); err != nil {
		c.JSON(http.StatusOK, &HTTPResp{
			ErrCode: errors.Type(err),
			ErrMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &HTTPResp{})
	return
}

// UserQueryResp 用户查询结果
type UserQueryResp struct {
	*HTTPResp
	User *model.UserInfo `json:"user"`
}

// UserQuery 用户创建
func UserQuery(c *gin.Context) {
	id := c.Query("id")
	var resp UserQueryResp

	user, err := userInfo.Query(context.Background(), id)
	if err != nil {
		resp.ErrCode = errors.Type(err)
		resp.ErrMsg = err.Error()
		c.JSON(http.StatusOK, &resp)
		return
	}

	resp.User = user

	c.JSON(http.StatusOK, &resp)
	return
}

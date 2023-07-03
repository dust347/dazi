package service

import (
	"context"
	"net/http"

	"github.com/dust347/dazi/internal/domain"
	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/errors"
	"github.com/dust347/dazi/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
)

var userInfo domain.UserInfoManager

// MiddlewareAuth 身份验证中间件
func MiddlewareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 header 中获取 token
		token := c.Request.Header.Get("token")
		// 验证
		id, ok := jwt.Parse(token)
		if !ok {
			c.Abort()
			c.JSON(http.StatusUnauthorized, HTTPResp{
				ErrCode: -1,
				ErrMsg:  "unauthorized",
			})
		}

		// 设置解析后的 user_id
		c.Request.Header.Set("user_id", id)

		c.Next()
	}
}

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

// LoginReq 登录请求
type LoginReq struct {
	JSCode string          `json:"js_code"`
	User   *model.UserInfo `json:"user"`
}

// LoginResp 登录返回
type LoginResp struct {
	HTTPResp
	User *model.UserInfo `json:"user"`
}

// Login 用户创建
func Login(c *gin.Context) {
	var req LoginReq
	var resp LoginResp

	if err := c.BindJSON(&req); err != nil {
		resp.ErrCode = errors.ParamErr
		resp.ErrMsg = "body unmarshal err"
		c.Abort()
		c.JSON(http.StatusInternalServerError, &resp)
		return
	}
	user, err := userInfo.Login(context.Background(), req.JSCode, req.User)

	if err != nil {
		resp.ErrCode = errors.Type(err)
		resp.ErrMsg = err.Error()
		c.Abort()
		c.JSON(http.StatusInternalServerError, &resp)
		return
	}

	// 创建 token
	token, err := jwt.Sign(user.ID)
	if err != nil {
		resp.ErrCode = errors.Type(err)
		resp.ErrMsg = err.Error()
		c.Abort()
		c.JSON(http.StatusInternalServerError, &resp)
	}

	c.Writer.Header().Set("token", token)
	resp.User = user
	c.JSON(http.StatusOK, &resp)
	return
}

// UserUpdate 用户创建
func UserUpdate(c *gin.Context) {
	var user model.UserInfo
	if err := c.BindJSON(&user); err != nil {
		c.Abort()
		c.JSON(http.StatusInternalServerError, &HTTPResp{
			ErrCode: errors.ParamErr,
			ErrMsg:  "body unmarshal err",
		})
		return
	}

	user.ID = c.Request.Header.Get("user_id")

	if err := userInfo.Update(context.Background(), &user); err != nil {
		c.Abort()
		c.JSON(http.StatusInternalServerError, &HTTPResp{
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
	HTTPResp
	User *model.UserInfo `json:"user"`
}

// UserQuery 用户查询
func UserQuery(c *gin.Context) {
	id := c.Request.Header.Get("user_id")
	var resp UserQueryResp

	user, err := userInfo.Query(context.Background(), id)
	if err != nil {
		resp.ErrCode = errors.Type(err)
		resp.ErrMsg = err.Error()
		c.Abort()
		c.JSON(http.StatusInternalServerError, &resp)
		return
	}

	resp.User = user

	c.JSON(http.StatusOK, &resp)
	return
}

// NearbyReq 附近的人请求
type NearbyReq struct {
	Location *model.Location `json:"location"`
}

// NearbyResp 附近的人返回
type NearbyResp struct {
	HTTPResp
	Users []model.UserInfo `json:"users"`
}

// Nearby 查找附近用户
func Nearby(c *gin.Context) {
	id := c.Request.Header.Get("user_id")
	var req NearbyReq
	var resp NearbyResp

	if err := c.BindJSON(&req); err != nil {
		c.Abort()
		c.JSON(http.StatusInternalServerError, &HTTPResp{
			ErrCode: errors.ParamErr,
			ErrMsg:  "body unmarshal err",
		})
		return
	}

	users, err := userInfo.Nearby(context.Background(), id, req.Location)
	if err != nil {
		resp.ErrCode = errors.Type(err)
		resp.ErrMsg = err.Error()
		c.Abort()
		c.JSON(http.StatusInternalServerError, &resp)
		return
	}

	resp.Users = users

	c.JSON(http.StatusOK, &resp)
	return
}

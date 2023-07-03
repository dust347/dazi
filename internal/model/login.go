package model

// LoginCheckReq 登录校验请求
type LoginCheckReq struct {
	Code string `json:"js_code"` // wx 登录时获取的 code
}

// LoginCheckResp 登录校验结果
type LoginCheckResp struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`

	SessionKey string `json:"session_key"` // 会话秘钥
	UniID      string `json:"unionid"`
	OpenID     string `json:"openid"`
}

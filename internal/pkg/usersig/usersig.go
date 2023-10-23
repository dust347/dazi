package usersig

import (
	"github.com/dust347/dazi/internal/pkg/config"
	"github.com/tencentyun/tls-sig-api-v2-golang/tencentyun"
)

// GenUserSig 生成 user sig
func GenUserSig(userid string, expire int) (string, error) {
	return tencentyun.GenUserSig(config.GetConfig().IM.AppID, config.GetConfig().IM.SecretKey, userid, expire)
}

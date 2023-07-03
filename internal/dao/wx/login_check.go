package wx

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/errors"
)

// LoginCheckClient 登录校验
type LoginCheckClient struct {
	prx *http.Client
	cfg *model.DatabaseConfig
}

// NewLoginCheckClient 创建 LoginCheckClient 实例
func NewLoginCheckClient(cfg *model.DatabaseConfig) (*LoginCheckClient, error) {
	if cfg == nil {
		return nil, errors.New(errors.ParamErr, "cfg is nil")
	}

	return &LoginCheckClient{
		prx: &http.Client{},
		cfg: cfg,
	}, nil
}

// Check 校验登录状态
func (cli *LoginCheckClient) Check(ctx context.Context, req *model.LoginCheckReq) (*model.LoginCheckResp, error) {
	if req == nil {
		return nil, errors.New(errors.ParamErr, "req is nil")
	}
	r, err := cli.prx.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		cli.cfg.Namespace, cli.cfg.Name, req.Code))
	if err != nil {
		return nil, errors.WithMsg(err, "http get err")
	}

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.WithMsg(err, "read resp body err")
	}

	var resp model.LoginCheckResp
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, errors.WithMsg(err, "body unmarshal err")
	}

	return &resp, nil
}

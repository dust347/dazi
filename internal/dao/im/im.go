package im

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/config"
	"github.com/dust347/dazi/internal/pkg/errors"
	"github.com/dust347/dazi/internal/pkg/usersig"
)

// Client im clientk
type Client struct {
	prx *http.Client

	appid      string
	identifier string
	key        string
}

// NewClient 创建 client 实例
func NewClient(cfg *model.DatabaseConfig) (*Client, error) {
	if cfg == nil {
		return nil, errors.New(errors.ParamErr, "cfg is nil")
	}

	return &Client{
		prx:        &http.Client{},
		identifier: cfg.Target,
	}, nil
}

// Account 账户信息
type Account struct {
	UserID  string `json:"UserID"`
	Nick    string `json:"Nick"`
	FaceURL string `json:"FaceUrl"`
}

// ImportAccount 导入用户至 im
func (cli *Client) ImportAccount(ctx context.Context, userID, nick, avatar string) error {
	account := Account{
		UserID:  userID,
		Nick:    nick,
		FaceURL: avatar,
	}
	b, err := json.Marshal(&account)
	if err != nil {
		return errors.WithMsg(err, "marshal err")
	}

	sig, err := usersig.GenUserSig(cli.identifier, 86400)

	resp, err := cli.prx.Post(
		fmt.Sprintf("https://%s/v4/im_open_login_svc/account_import?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=json",
			"console.tim.qq.com", config.GetConfig().IM.AppID, cli.identifier, sig, rand.Int31()),
		"application/json",
		bytes.NewReader(b),
	)
	if err != nil {
		return errors.WithMsg(err, "post err")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.Errorf(errors.UnknownErr, "post err, status: %s", resp.Status)
	}

	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return errors.WithMsg(err, "read body err")
	}
	log.Printf("im resp: %s", b)

	var imResp Resp
	if err := json.Unmarshal(b, &imResp); err != nil {
		return errors.WithMsg(err, "unmarshal err")
	}

	if imResp.ErrorCode != 0 {
		return errors.New(imResp.ErrorCode, imResp.ErrorInfo)
	}

	return nil
}

// Resp im 返回的resp
type Resp struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int32
}

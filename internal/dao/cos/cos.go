package cos

import (
	"context"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// Client cos client
type Client struct {
	url string
	prx *cos.Client
}

// NewClient 创建 Client 实例
func NewClient(cfg *model.DatabaseConfig) (*Client, error) {
	if cfg == nil {
		return nil, errors.New(errors.ParamErr, "cfg is nil")
	}

	u, err := url.Parse(cfg.Target)
	if err != nil {

	}
	b := &cos.BaseURL{BucketURL: u}
	prx := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: cfg.Namespace, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: cfg.Name, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	return &Client{
		url: cfg.Target,
		prx: prx,
	}, nil
}

// Upload 上次图片
func (cli *Client) Upload(ctx context.Context, fileName string, image model.ImageFile) (string, error) {
	_, err := cli.prx.Object.Put(ctx, fileName, image, nil)
	if err != nil {
		return "", errors.WithMsg(err, "upload err")
	}

	return filepath.Join(cli.url, fileName), nil
}

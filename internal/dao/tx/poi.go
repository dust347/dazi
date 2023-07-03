package tx

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/dust347/dazi/internal/model"
	"github.com/dust347/dazi/internal/pkg/errors"
)

// PoiClient poi client
type PoiClient struct {
	prx *http.Client
	cfg *model.DatabaseConfig
}

// NewPoiClient 创建 PoiClient 实例
func NewPoiClient(cfg *model.DatabaseConfig) (*PoiClient, error) {
	if cfg == nil {
		return nil, errors.New(errors.ParamErr, "cfg is nil")
	}

	return &PoiClient{
		prx: &http.Client{},
		cfg: cfg,
	}, nil
}

const path = "/ws/geocoder/v1/"

func (cli *PoiClient) sign(loc *model.Location) string {
	m := map[string]string{
		"key":      cli.cfg.Namespace,
		"location": loc.String(),
	}
	keys := []string{"key", "location"}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	b := &strings.Builder{}
	b.WriteString(path)
	b.WriteString("?")
	for i, k := range keys {
		if i > 0 {
			b.WriteString("&")
		}
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(m[k])
	}
	b.WriteString(cli.cfg.Target)

	h := md5.New()
	h.Write([]byte(b.String()))
	return hex.EncodeToString(h.Sum(nil))
}

// GetCity 获取城市信息
func (cli *PoiClient) GetCity(ctx context.Context, loc *model.Location) (*model.CityInfo, error) {
	if loc == nil {
		return nil, errors.New(errors.ParamErr, "location is nil")
	}

	resp, err := cli.prx.Get(fmt.Sprintf("https://apis.map.qq.com%s?key=%s&location=%s&sig=%s",
		path, cli.cfg.Namespace, loc.String(), cli.sign(loc)))
	if err != nil {
		return nil, errors.WithMsg(err, "get city info err")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf(errors.UnknownErr, "status: %s", resp.Status)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	var poi PoiResp
	if err := json.Unmarshal(b, &poi); err != nil {
		return nil, errors.WithMsg(err, "unmarshal err")
	}

	return &model.CityInfo{
		CityCode: poi.Result.AdInfo.CityCode,
		CityName: poi.Result.AdInfo.CityName,
	}, nil
}

// PoiResp 返回结果
type PoiResp struct {
	Status int32  `json:"status"`
	Msg    string `json:"message"`
	Result Result `json:"result"`
}

// Result ...
type Result struct {
	AdInfo AdInfo `json:"ad_info"`
}

// AdInfo ...
type AdInfo struct {
	CityCode string `json:"city_code"`
	CityName string `json:"city"`
}

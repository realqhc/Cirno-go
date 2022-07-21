package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/imroc/req"
	"github.com/zsakvo/Cirno-go/config"
)

var client *req.Req

func Get(url string, paras req.Param) ([]byte, error) {
	cfg := config.Config
	var err error
	var r *req.Resp
	var res []byte
	var param req.Param
	if url == "/signup/login" {
		param = req.Param{
			"app_version": cfg.App.AppVersion,
		}
	} else {
		param = req.Param{
			"account":     cfg.App.Account,
			"app_version": cfg.App.AppVersion,
			"login_token": cfg.App.LoginToken,
		}
	}
	for k, v := range paras {
		param[k] = v
	}
	if client == nil {
		client = req.New()
	}
	client.SetTimeout(20 * time.Second)
	r, err = client.Get(cfg.App.HostUrl+url, param, req.Header{"User-Agent": cfg.App.UserAgent})
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	res, err = Decode(r.String(), cfg.App.DefaultKey)
	if !strings.Contains(string(res), "100000") {
		err = errors.New(ErrTip(string(res)))
	}
	return res, err
}

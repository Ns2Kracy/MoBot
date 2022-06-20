package util

import (
	"MoBot/global"
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(url string) *http.Response {
	defer func() {
		info := recover()
		if info != nil {
			global.GVA_LOG.Info("recover from http.Get", zap.Any("info", info))
		}
	}()
	rsp, err := http.Get(url)
	if err != nil {
		global.GVA_LOG.Error("http.Get", zap.Any("err", err))
	}
	return rsp
}

func HttpProxyGet(URL string, proxy string) *http.Response {
	uri, err := url.Parse(proxy)

	client := http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(uri),
		},
	}
	defer func() {
		info := recover()
		if info != nil {
			global.GVA_LOG.Info("recover from http.Get", zap.Any("info", info))
		}
	}()

	rsp, err := client.Get(URL)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	return rsp
}
func HttpPostJson(url string, data interface{}) *http.Response {
	defer func() {
		info := recover()
		if info != nil {
			global.GVA_LOG.Info("recover from http.Post", zap.Any("info", info))
		}
	}()
	jsonStr, _ := json.Marshal(data)
	rsp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		global.GVA_LOG.Error("http.Post", zap.Any("err", err))
	}
	return rsp
}

func HttpPostForm(url string, body url.Values) *http.Response {
	defer func() {
		info := recover()
		if info != nil {
			global.GVA_LOG.Info("recover from http.PostForm", zap.Any("info", info))
		}
	}()
	rsp, err := http.PostForm(url, body)
	rsp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rsp.Header.Set("Accept", "application/x-www-form-urlencoded")
	if err != nil {
		global.GVA_LOG.Error("http.PostForm", zap.Any("err", err))
	}
	// 延迟关闭(十分重要)
	defer rsp.Body.Close()
	return rsp
}

func GetRspBody(rsp *http.Response) []byte {
	body, _ := ioutil.ReadAll(rsp.Body)
	return body
}

// 一个通用的请求方法, 并设定请求格式是json还是form, 如果不是post请求, 则不需要设定
func NewHttpRequest(method, url string, body, contentType interface{}) *http.Request {
	var req *http.Request
	var err error
	if method == "POST" {
		if contentType == global.JSON_Type {
			req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(body.(string))))
		} else if contentType == global.Form_Type {
			req, err = http.NewRequest(method, url, bytes.NewBufferString(body.(string)))
		} else {
			req, err = http.NewRequest(method, url, nil)
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		global.GVA_LOG.Error("http.NewRequest", zap.Any("err", err))
	}
	return req
}

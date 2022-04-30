package util

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(url string) ([]byte, error) {
	defer func() {
		info := recover()
		if info != nil {
			iris.New().Logger().Info("recover from http.Get", info)
		}
	}()
	rsp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	return GetRspBody(rsp), nil
}

func HttpPostJson(url string, data interface{}) ([]byte, error) {
	defer func() {
		info := recover()
		if info != nil {
			iris.New().Logger().Info("recover from http.Post", info)
		}
	}()
	jsonStr, _ := json.Marshal(data)
	rsp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	return GetRspBody(rsp), nil
}

func HttpPostForm(url string, body url.Values) ([]byte, error) {
	defer func() {
		info := recover()
		if info != nil {
			iris.New().Logger().Info("recover from http.PostFor", info)
		}
	}()
	rsp, err := http.PostForm(url, body)
	rsp.Header.Set("Content-Type", "application/json")
	rsp.Header.Set("Accept", "application/json")
	if err != nil {
		panic(err)
	}
	// 延迟关闭(十分重要)
	defer rsp.Body.Close()
	return GetRspBody(rsp), nil
}

func GetRspBody(rsp *http.Response) []byte {
	body, _ := ioutil.ReadAll(rsp.Body)
	return body
}

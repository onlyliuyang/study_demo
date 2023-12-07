package demo

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/faabiosr/cachego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type AccessTokenManager interface {
	GetName() (name string)
	GetAccessToken() (accessToken string, err error)
}

type getRefreshRequestFunc func() *http.Request

type DefaultAccessTokenManager struct {
	Id                    string
	Name                  string
	GetRefreshRequestFunc getRefreshRequestFunc
	Cache                 cachego.Cache
}

// 防止多个 goroutine 并发刷新冲突
var getAccessTokenLock sync.Mutex

// GetAccessToken 获取access_token
func (m *DefaultAccessTokenManager) GetAccessToken() (accessToken string, err error) {
	accessToken = m.getAccessTokenFromChannelPlatform()
	return accessToken, nil

	cacheKey := m.getCacheKey()
	accessToken, err = m.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	getAccessTokenLock.Lock()
	defer getAccessTokenLock.Unlock()

	accessToken, err = m.Cache.Fetch(cacheKey)
	if accessToken != "" {
		return
	}

	req := m.GetRefreshRequestFunc()
	// 添加 serverUrl
	if !strings.HasPrefix(req.URL.String(), "http") {
		parse, _ := url.Parse(ServerUrl)
		req.URL.Host = parse.Host
		req.URL.Scheme = parse.Scheme
	}

	req.Header.Set("Content-Type", contentTypeApplicationJson)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var result = struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			AccessToken string  `json:"access_token"`
			ExpiresIn   float64 `json:"expires_in"`
		} `json:"data"`
	}{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		err = fmt.Errorf("unmarshal error %s", string(resp))
		return
	}

	if result.Data.AccessToken == "" {
		err = fmt.Errorf("%s", string(resp))
		return
	}

	accessToken = result.Data.AccessToken

	err = m.Cache.Save(cacheKey, accessToken, time.Duration(result.Data.ExpiresIn)*time.Second)
	if err != nil {
		return
	}

	return
}

// getCacheKey
func (m *DefaultAccessTokenManager) getCacheKey() (key string) {
	return "access_token:" + m.Id
}

// GetName 获取 access_token 参数名称
func (m *DefaultAccessTokenManager) GetName() (name string) {
	return m.Name
}

func (m *DefaultAccessTokenManager) getAccessTokenFromChannelPlatform() string {
	api := "https://platform.douyuxingchen.com/api/xiaoetong/get_access_token"
	params := url.Values{}

	params.Set("app_id", "appby7f3kfu2684")
	//params.Set("refresh", "")

	//fmt.Println(params.Encode() + "&key=CHANNELPLATFORM")

	sign := MD5Bytes([]byte(params.Encode() + "&key=CHANNELPLATFORM"))
	params.Set("sign", sign)
	response, _ := http.Get(api + "?" + params.Encode())

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	type Result struct {
		Code int               `json:"code"`
		Data map[string]string `json:"data"`
	}
	var result Result
	//fmt.Println(string(resp))
	err = json.Unmarshal(resp, &result)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result.Data["token"]

}

func MD5Bytes(s []byte) string {
	ret := md5.Sum(s)
	return hex.EncodeToString(ret[:])
}

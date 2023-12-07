package main

import (
	"encoding/json"
	"fmt"
	"github.com/faabiosr/cachego/file"
	"github.com/testProject/xiaoetong/demo"
	"net/http"
	"os"
	"strings"
)

const (

	//豆语星辰
	XiaoEAppId     = "appby7f3kfu2684"                  // 店铺app_id
	XiaoEClientId  = "xopjUYBBebG7514"                  // 店铺client_id
	XiaoEAppSecret = "x5nocBxeXlKkv9w0RmTi4PEuBsGmgoYg" // 店铺client_sercet
	XiaoEGrantType = "client_credential"                // 固定不变

	//豆神网校
	//XiaoEAppId     = "appby7f3kfu2684"                  // 店铺app_id
	//XiaoEClientId  = "xopjUYBBebG7514"                  // 店铺client_id
	//XiaoEAppSecret = "x5nocBxeXlKkv9w0RmTi4PEuBsGmgoYg" // 店铺client_sercet
	//XiaoEGrantType = "client_credential"                // 固定不变

	GetUserListOpenApi     = "https://api.xiaoe-tech.com/xe.user.batch.get/1.0.0"
	GetUserStudyOpenAPi    = "https://api.xiaoe-tech.com/xe.user.learning.records.get/1.0.0"
	ResigerUserOpenAPI     = "https://api.xiaoe-tech.com/xe.user.register/1.0.0"
	GetLiveUserListOpenAPI = "https://api.xiaoe-tech.com/xe.alive.user.list/1.0.0"
	GetLiveDetailOpenAPI   = "https://api.xiaoe-tech.com/xe.alive.detail.get/1.0.0"
	GetLiveListOpenAPI     = "https://api.xiaoe-tech.com/xe.alive.list.get/1.0.0"
	GetUserDetailOpenAPI   = "https://api.xiaoe-tech.com/xe.user.info.get/1.0.0"
)

func main() {
	//registerUser()
	//getUserStudyData()

	//getLiveUserData()
	//getLiveDetail()
	//getLiveList()
	getUserDetail()
}

func getClient() *demo.Client {
	xiaoE := &demo.DefaultAccessTokenManager{
		Id:   XiaoEAppId,
		Name: "access_token",
		GetRefreshRequestFunc: func() *http.Request {
			params := make(map[string]string)
			params["app_id"] = XiaoEAppId
			params["client_id"] = XiaoEClientId
			params["secret_key"] = XiaoEAppSecret
			params["grant_type"] = XiaoEGrantType

			str, err := json.Marshal(params)
			if err != nil {
				fmt.Println(err)
			}
			payload := strings.NewReader(string(str))
			req, err := http.NewRequest(http.MethodGet, demo.ServerUrl, payload)
			if err != nil {
				return nil
			}

			return req
		},
		// os.TempDir() - 设置缓存路径，默认为系统默认缓存路径,为了方便查找建议修改路径
		Cache: file.New(os.TempDir()),
	}

	// 小鹅云 客户端
	xiaoEClient := demo.NewClient(xiaoE)
	return xiaoEClient
}

func getUserList() {
	xiaoEClient := getClient()

	// 调用示例 不需要传入access_token
	UserListParams := make(map[string]interface{})
	UserListParams["page"] = 1
	UserListParams["page_size"] = 2

	resp, err := xiaoEClient.CurlDo(http.MethodPost, GetUserListOpenApi, UserListParams)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(string(resp))
}

func getUserStudyData() {
	xiaoEClient := getClient()

	// 调用示例 不需要传入access_token
	user_ids := []string{
		"u_api_656943862ba39_qAxgIExEEb",
	}

	for _, userId := range user_ids {
		fmt.Println(userId)
		UserStudyParams := make(map[string]interface{})
		UserStudyParams["user_id"] = userId
		UserStudyParams["data"] = map[string]int{
			"page":      1,
			"page_size": 10,
		}

		resp, err := xiaoEClient.CurlDo(http.MethodPost, GetUserStudyOpenAPi, UserStudyParams)
		if err != nil {
			fmt.Println("err: ", err)
		}
		fmt.Println(string(resp))
	}
}

func registerUser() {
	client := getClient()
	params := make(map[string]interface{})
	params["data"] = map[string]string{
		"phone": "15962193629",
	}
	resp, err := client.CurlDo(http.MethodPost, ResigerUserOpenAPI, params)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(string(resp))
}

func getLiveUserData() {
	client := getClient()
	params := make(map[string]interface{})
	params["resource_id"] = "l_6556e771e4b04c103862be1a"
	params["key_word_type"] = 4 //用户id
	params["key_word"] = "u_655717865650f_faZTZyFrtN"
	params["page"] = 1
	params["page_size"] = 10

	resp, err := client.CurlDo(http.MethodPost, GetLiveUserListOpenAPI, params)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(string(resp))
}

func getLiveDetail() {
	client := getClient()
	params := make(map[string]interface{})
	params["id"] = "l_6556e771e4b04c103862be1a"

	resp, err := client.CurlDo(http.MethodPost, GetLiveDetailOpenAPI, params)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(string(resp))
}

func getLiveList() {
	client := getClient()
	params := make(map[string]interface{})
	params["search_content"] = "语文各阶段学习规划"

	resp, err := client.CurlDo(http.MethodPost, GetLiveListOpenAPI, params)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(string(resp))
}

func getUserDetail() {
	client := getClient()
	params := make(map[string]interface{})
	params["user_id"] = "u_655b3f4e61c28_WTlDJyBWGq"
	params["data"] = map[string]interface{}{
		"field_list": []string{
			"wx_union_id",
			"wx_open_id",
			"wx_app_open_id",
			"wx_email",
			"phone_collect",
		},
	}

	resp, err := client.CurlDo(http.MethodPost, GetUserDetailOpenAPI, params)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(string(resp))
}

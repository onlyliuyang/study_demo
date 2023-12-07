package main

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	URL   = "https://api.ip138.com/mobile/"
	TOKEN = "3474f1c6910171d3aa40e4852a0eafd0"
)

//----------------------------------
// 手机号查询接口调用示例代码
//----------------------------------

// xml struct
type xmlinfo struct {
	Ret    string          `xml:"ret"`
	Mobile string          `xml:"mobile"`
	Data   locationxmlInfo `xml:"data"`
}

type locationxmlInfo struct {
	Province string `xml:"province"`
	City     string `xml:"city"`
	Card     string `xml:"card"`
	Zone     string `xml:"zone"`
}

// json struct
type jsoninfo struct {
	Ret    string    `json:"ret"`
	Mobile string    `json:"mobile"`
	Data   [4]string `json:"data"`
}

func main() {
	//mobileLocation("15176082486", "xml")

	payload := "eyJhcHBfa2V5IjoiMjc1NjY4YmE2NTUwNDljZDczOWQxZDllNmIzMWNjZjEiLCJhcHBfc2VjcmV0IjoiN2M5NzI2NjMxNzBkNmJjMTg0ODRkMDViYzk4NzIyZjQiLCJleHAiOjE2ODM2MjczNjIsImlzcyI6ImJsb2ctc2VydmljZSJ9"
	p, _ := base64.StdEncoding.DecodeString(payload)
	fmt.Println(string(p))
}

func mobileLocation(mobile string, dataType string) {

	queryUrl := fmt.Sprintf("%s?mobile=%s&datatype=%s", URL, mobile, dataType)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", queryUrl, nil)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	reqest.Header.Add("token", TOKEN)
	response, err := client.Do(reqest)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)

		if dataType == "jsonp" {
			var info jsoninfo
			json.Unmarshal(bodyByte, &info)
			fmt.Println(info.Mobile)
		} else if dataType == "xml" {
			var info xmlinfo
			xml.Unmarshal(bodyByte, &info)
			fmt.Println(info.Mobile)
		}

		//body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(bodyByte))
	}

	return
}

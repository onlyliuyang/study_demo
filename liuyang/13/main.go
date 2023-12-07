package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type LogData struct {
	IndexStatus string `json:"__INDEX_STATUS__"`
	RawLog      string `json:"__RAWLOG__"`
}

type PhaseData struct {
	PhaseClass    string      `json:"phase_class"`
	PhaseFunction string      `json:"phase_function"`
	PhaseLine     int         `json:"phase_line"`
	PhaseLevel    string      `json:"phase_level"`
	PhaseMsg      string      `json:"phase_msg"`
	PhaseTime     int64       `json:"phase_time"`
	PhaseCost     string      `json:"phase_cost"`
	PhaseContent  interface{} `json:"phase_content"`
	// ... 其他字段
}

type PhaseContent struct {
	Authorization string `json:"authorization"`
	Token         string `json:"token"`
	// ... 其他字段
}

type Content struct {
	ReqUrl  string `json:"req_url"`
	ClickId string `json:"click_id"`
	Url     string `json:"url"`
}

type RawLogData struct {
	FileName  string `json:"__FILENAME__"`
	HostName  string `json:"__HOSTNAME__"`
	Source    string `json:"__SOURCE__"`
	AppName   string `json:"app_name"`
	Cost      string `json:"cost"`
	Level     string `json:"level"`
	LocalIP   string `json:"local_ip"`
	LocalPort string `json:"local_port"`
	LogID     string `json:"log_id"`
	LogTag    string `json:"log_tag"`
	Message   string `json:"message"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	PhaseData string `json:"phase_data"`
	ReqTime   string `json:"req_time"`
	Request   string `json:"request"`
	ResTime   string `json:"res_time"`
	Response  string `json:"response"`
	Status    string `json:"status"`
	Time      string `json:"time"`
	UserID    string `json:"user_id"`
	UserIP    string `json:"user_ip"`
	// ... 其他字段
}

const FILE_DIR = "/Users/admin/Downloads/log_100020398441_2938139c-ebf2-4f93-aab7-f193bed88037_20231116_export-a54be036-508a-472e-b02e-4b643c1c3b79_1700123635/export-a54be036-508a-472e-b02e-4b643c1c3b79.json"

func main() {
	readFile(FILE_DIR)
}

func readFile(filePath string) {
	fileHeader, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHeader.Close()

	bytes, err := io.ReadAll(fileHeader)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	list := strings.Split(string(bytes), "\n")
	var clickIdList []string
	var clickIdMap map[string]bool
	clickIdMap = make(map[string]bool)
	for _, item := range list {

		if len(item) <= 0 {
			continue
		}
		// 解析JSON数据到LogData结构体
		var logData LogData
		err := json.Unmarshal([]byte(item), &logData)
		if err != nil {
			fmt.Println(item)
			fmt.Println("Error decoding JSON1:", err)
			return
		}

		// 解析嵌套的RawLog JSON数据到RawLogData结构体
		var rawLogData RawLogData
		err = json.Unmarshal([]byte(logData.RawLog), &rawLogData)
		if err != nil {
			fmt.Println("Error decoding RawLog JSON2:", err)
			return
		}

		var phdataList []PhaseData
		err = json.Unmarshal([]byte(rawLogData.PhaseData), &phdataList)
		if err != nil {
			fmt.Println("Error decoding RawLog JSON3:", err)
			return
		}

		adInfo := phdataList[2].PhaseContent.(map[string]interface{})
		if clickId, ok := adInfo["click_id"]; ok {
			if _, ok := clickIdMap[clickId.(string)]; !ok {
				clickIdList = append(clickIdList, "\""+clickId.(string)+"\"")
				clickIdMap[clickId.(string)] = true
				count++
			}
		}
	}
	fmt.Println(count)
	fmt.Println(strings.Join(clickIdList, ","))

}

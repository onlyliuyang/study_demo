package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func readline(filePath string) []string {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	var lines []string
	var i int = 0
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ",")
		lines = append(lines, splits[1])
		if i > 10 {
			//break
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, err)
	}
	return lines
}

//实现单个文件下载
func download(url string, fileId string) string {
	filePath := "download/"
	newFile, err := os.Create(filePath + fileId)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer newFile.Close()

	client := http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer resp.Body.Close()

	_, err = io.Copy(newFile, resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return fileId
}

func main() {
	fileList := readline("./urlmaps_凯叔生产_题库数据映射.csv")
	//fmt.Println(fileList)

	//ch := make(chan string)

	for _, url := range fileList {
		urlSplit := strings.Split(url, "/")
		fileId := urlSplit[len(urlSplit)-1]
		//go func(url, filed string) {
		//	ch <- download(url, filed)
		//}(url, fileId)
		res := download(url, fileId)
		fmt.Println(res, "success")
	}
	//
	//timeout := time.After(1 * time.Hour)
	//for i := 0; i < len(fileList); i++ {
	//	select {
	//	case res := <-ch:
	//		fmt.Println(res, "success")
	//	case <-timeout:
	//		fmt.Println("timeout")
	//	}
	//}
}

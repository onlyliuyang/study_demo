package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	HOST      = "https://test-mapi.douyuxingchen.com/ai/userChats/aiAsk"
	TOKEN     = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImp0aSI6ImRvdXl1eGluZ2NoZW4ifQ.eyJpc3MiOiJodHRwOlwvXC9sb2NhbGhvc3QiLCJhdWQiOiJodHRwOlwvXC9sb2NhbGhvc3QiLCJqdGkiOiJkb3V5dXhpbmdjaGVuIiwiaWF0IjoxNjgxMjg0MjUyLCJleHAiOjE3MTI4MjAyNTIsInVpZCI6MTI3NiwicHJvZHVjdCI6MSwicGxhdGZvcm0iOjN9.a-F-58xXqNLej1gZaEEmKyskFqLc9qE_WiVWPtirpDc"
	questions = []string{
		"你好",
		"你是谁",
		"你吃饭了吗",
		"你多大了",
		"你几岁了",
		"你上学了吗",
		"你明天去上学吗",
		"你家在哪里",
		"你是什么学历",
		"你高中在哪里上的",
		"你是哪里人呢",
		"你会写代码吗",
		"请打印出π的第3000位值",
		"讲一下牛顿第二定律",
		"地球是圆的吗",
		"什么时候解放台湾",
		"台湾有茶叶蛋吗",
		"航母上边有多少条螺丝",
		"你们有加班费吗",
		"1+1在什么情况下不等于2",
	}
)

type Body struct {
	CourseId             int    `json:"course_id"`
	LessonId             int    `json:"lesson_id"`
	LessonTaskResourceId int    `json:"lesson_task_resource_id"`
	KsExercisesId        int    `json:"ks_exercises_id"`
	PlayTime             int    `json:"play_time"`
	EventType            int    `json:"event_type"`
	Content              string `json:"content"`
	OptionId             int    `json:"option_id"`
	LessonTaskId         int    `json:"lesson_task_id"`
	ClientUnique         string `json:"clientUnique"`
}

func main() {
	qLen := len(questions)
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 1; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				rand.Seed(time.Now().UnixNano())
				randInt := rand.Intn(qLen)
				question := questions[randInt]
				body := Body{
					CourseId:             523,
					LessonId:             906,
					LessonTaskResourceId: 987,
					KsExercisesId:        8,
					PlayTime:             20,
					EventType:            1,
					Content:              question,
					OptionId:             1,
					LessonTaskId:         10,
					ClientUnique:         "test",
				}
				resp := POST(body)
				fmt.Println(resp)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Done")
}

func POST(data Body) (r string) {
	var client *http.Client
	var request *http.Request
	var resp *http.Response
	var body []byte

	client = &http.Client{Timeout: time.Second * 2}
	bodyMarshal, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bodyMarshal))
	reqBody := strings.NewReader(string(bodyMarshal))
	request, err = http.NewRequest(http.MethodPost, HOST, reqBody)
	request.Header.Set("Authorization", TOKEN)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	return string(body)
}

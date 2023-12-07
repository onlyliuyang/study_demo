package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
后端机试题目

姓名：                日期：


题目：公司有n个组，每组人数相同，>=1人，需要进行随机的组队吃饭。
要求：1. 两两一队，不能落单，落单则三人一队
  2. 一个人只出现一次
  3. 队伍中至少包含两个组
  4. 随机组队，重复执行程序得到的结果不一样
举例：
GroupList = [  # 小组列表
    ['小名', '小红', '小马', '小丽', '小强'],
    ['大壮', '大力', '大1', '大2', '大3'],
    ['阿花', '阿朵', '阿蓝', '阿紫', '阿红'],
    ['A', 'B', 'C', 'D', 'E'],
    ['一', '二', '三', '四', '五'],
    ['建国', '建军', '建民', '建超', '建跃'],
    ['爱民', '爱军', '爱国', '爱辉', '爱月']
]

输入：GroupList
输出：(A, 小名)，（B, 小红）。。。


*/

func main() {
	groupList := [][]string{
		{"小名", "小红", "小马", "小丽", "小强"},
		{"大壮", "大力", "大1", "大2", "大3"},
		{"阿花", "阿朵", "阿蓝", "阿紫", "阿红"},
		{"A", "B", "C", "D", "E"},
		{"一", "二", "三", "四", "五"},
		{"建国", "建军", "建民", "建超", "建跃"},
		{"爱民", "爱军", "爱国", "爱辉", "爱月"},
	}

	teamWay := getTeamWay(groupList)
	fmt.Println(teamWay)
}

func getTeamWay(groupList [][]string) [][]string {
	result := [][]string{[]string{}}
	groupLength := len(groupList)
	memberLength := len(groupList[0])

	beginGroup := 0
	lastGroup := -1
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for beginGroup < groupLength {
		randG := r.Intn(groupLength) //随机选择一个队开始
		if lastGroup == randG {
			continue
		}
		randM := r.Intn(memberLength) //随机选择一个成员
		if len(groupList[randG]) == 0 || len(groupList[randG]) <= randM {
			continue
		}
		val := groupList[randG][randM]
		groupList[randG] = append(groupList[randG][:randM], groupList[randG][randM+1:]...) //移除已经组队的成员
		result[len(result)-1] = append(result[len(result)-1], val)

		lastGroup = randG
		if len(result[len(result)-1]) >= 2 {
			lastGroup = -1
			result = append(result, []string{})
		}
		if len(groupList[randG]) == 0 {
			beginGroup++
		}
	}

	//
	if len(result[len(result)-1]) == 1 {
		temp := result[len(result)-1]
		result = result[:len(result)-1]
		result[len(result)-1] = append(result[len(result)-1], temp...)
	}
	if len(result[len(result)-1]) == 0 {
		return result[:len(result)-1]
	}
	return result
}

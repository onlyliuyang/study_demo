package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var ALLOW_ACTIONS []string = []string{"checkout", "push", "pull", "merge", ""}

var (
	CHECKOUT_CMD        = "git checkout %s"
	PULL_CMD            = "git pull"
	CURRENT_FEATURE_CMD = "git branch --show-current"
)

var (
	/* 从哪个支会拉取 */
	FROME_FEATURE = ""
	/* 拉取出的分支 */
	CHECKOUT_FEATURE = ""
	/* 提交的备注 */
	PUSH_COMMENT = ""
	/* 准备合并的分支 */
	MERGE_TO_FEATURE = ""
	/* 动作 */
	ACTION = ""
)

func main() {
	flag.StringVar(&FROME_FEATURE, "f", "master", "拉取的源分支")
	flag.StringVar(&CHECKOUT_FEATURE, "c", "", "检出的分支")
	flag.StringVar(&PUSH_COMMENT, "m", "", "提交备注")
	flag.StringVar(&ACTION, "a", "", "执行的动作")
	flag.Parse()

	if ACTION == "" {
		fmt.Println("必须要输入执行的动作")
		return
	}

	switch ACTION {
	case "checkout":
		checkoutFeature(FROME_FEATURE, CHECKOUT_FEATURE)
	}
}

func checkoutFeature(fromFeature, toFeature string) {
	out, err := exec.Command("base", "-c", fmt.Sprintf(CHECKOUT_CMD, fromFeature)).Output()
	if err != nil {
		fmt.Println("切换 Master 分支发生错误: ", err.Error())
		return
	}
	fmt.Println(out)
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func WrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func (err MyError) Error() string {
	return err.Message
}

//lowLevel 模块
type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{WrapError(err, err.Error())}
		//return false, err
	}
	return info.Mode().Perm()&0100 == 0100, nil
}

//intermediate模块
type IntermediateErr struct {
	error
}

func runJob(id string) error {
	const jobBinPath = "/Users/admin/go/src/github.com/testProject/liuyang/05/03.go"
	//const jobBinPath = "/bin/ls"
	isExecutable, err := isGloballyExec(jobBinPath)
	//os.Exit(2)
	if err != nil {
		//return err
		return IntermediateErr{WrapError(err, "cannot run job %q: requisite binaries not available", id)}
	} else if isExecutable == false {
		return WrapError(nil, "cannot run job %q : requisite binaries are not excutable")
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logId: %v]:", key))
	log.Printf("%#v", err)
	log.Printf("[%v] %v", key, message)
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	err := runJob("1")
	if err != nil {
		msg := "There was an unexpected issue; please report this as a bug."
		if _, ok := err.(IntermediateErr); ok {
			msg = err.Error()
		}
		handleError(1, err, msg)
	}
}

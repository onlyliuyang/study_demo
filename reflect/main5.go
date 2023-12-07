package main

import "fmt"

var _ Study = (*study)(nil)

type Study interface {
	Listen(message string) string
}

type study struct {
	Name string
}

func (t *study) Listen(message string) string {
	return message
}

func Print(s Study) string {
	return s.Listen("kwg kwg")
}

func main() {
	fmt.Println("hello world")

	var stu study
	fmt.Println(stu.Listen("xiaoxiao"))

	fmt.Println(Print(&stu))
}

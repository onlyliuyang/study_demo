package main

import (
	"bytes"
	"fmt"
	"strings"
)

var base string = "123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASFGHJKLZXCVBNM"

func SumString(str string) string {
	return str + base
}

func SprintfString(str string) string {
	return fmt.Sprintf("%s%s", str, base)
}

func BuilderString(str string) string {
	var builder strings.Builder
	builder.Grow(2 * len(str))
	builder.WriteString(base)
	builder.WriteString(str)
	return builder.String()
}

func ByteString(str string) string {
	buf := new(bytes.Buffer)
	buf.WriteString(base)
	buf.WriteString(str)
	return buf.String()
}

func ByteSliceString(str string) string {
	buf := make([]byte, 0)
	buf = append(buf, base...)
	buf = append(buf, str...)
	return string(buf)
}

func JoinString(str string) string {
	return strings.Join([]string{base, str}, "")
}

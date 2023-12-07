package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Buff struct {
	Buffer *bytes.Buffer
	Writer *bufio.Writer
}

//初始化
func NewBuff() *Buff  {
	b := bytes.NewBuffer([]byte{})
	return &Buff{
		Buffer:b,
		Writer:bufio.NewWriter(b),
	}
}

func (b *Buff) WriteString(str string) error {
	_, err := b.Writer.WriteString(str)
	return err
}

func (b *Buff) SaveAs(name string) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := b.Writer.Flush(); err != nil {
		return err
	}

	_, err = b.Buffer.WriteTo(file)
	return err
}

func main()  {
	var b = NewBuff()
	err := b.WriteString("hash")
	fmt.Println(err)
}
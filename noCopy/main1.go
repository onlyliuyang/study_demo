package main

import "fmt"

type noCopy struct {

}

func (*noCopy) Lock() {

}

func (*noCopy) Unlock() {

}

type DoNotCopy struct {
	noCopy
}

func main()  {
	var d DoNotCopy
	d2 := d
	fmt.Println(d2)
}

package main

import (
	"errcodedemo/errcode"
	"fmt"
)

func main() {
	fmt.Println(errcode.ErrUnknown)
	fmt.Println(errcode.ErrNotFound)
}

package main

import (
	"fmt"
	"stringerdemo/errcode"
)

func main() {
	fmt.Println(errcode.ErrUnknown)
	fmt.Println(errcode.ErrNotFound)
}

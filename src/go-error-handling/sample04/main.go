package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	ERROR_MSG_01 = "doSomething is error. b is true"
	ERROR_MSG_02 = "doSomething is error. b is false"
)

var (
	ERROR_01 = errors.New(ERROR_MSG_01)
	ERROR_02 = errors.New(ERROR_MSG_02)
)

func doSomething(b bool) error {

	if b {
		return ERROR_01

	} else {
		return ERROR_02
	}

	return nil
}

func main() {

	err := doSomething(false)

	if err == ERROR_01 {
		fmt.Println("ERROR: インスタンスERROR_01に応じた処理を行う")
		os.Exit(1)

	} else if err == ERROR_02 {
		fmt.Println("ERROR: インスタンスERROR_02に応じた処理を行う")
		os.Exit(1)
	}
	fmt.Println("main is success")
}

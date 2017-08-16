package main

import (
	"fmt"
	"os"
)

const (
	ERROR_MSG_01 = "doSomething is error. b is true"
	ERROR_MSG_02 = "doSomething is error. b is false"
)

type MyError struct {
	Msg  string
	Code int
}

func (err *MyError) Error() string {
	return fmt.Sprintf("ERROR: %d %s", err.Code, err.Msg)
}

var (
	MyError_01 = &MyError{Msg: "this is MyError_001", Code: 30001}
	MyError_02 = &MyError{Msg: "this is MyError_002", Code: 30002}
)

func doSomething(b bool) error {

	if b {
		return MyError_01

	} else {
		return MyError_02
	}

	return nil
}

func main() {

	err := doSomething(false)

	if err == MyError_01 {
		fmt.Println("ERROR: インスタンスMyError_01に応じた処理を行う")
		os.Exit(1)

	} else if err == MyError_02 {
		fmt.Println("ERROR: インスタンスMyError_02に応じた処理を行う")
		os.Exit(1)
	}
	fmt.Println("main is success")
}

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

func doSomething(b bool) error {

	if b {
		return errors.New(ERROR_MSG_01)

	} else {
		return errors.New(ERROR_MSG_02)
	}

	return nil
}

func main() {

	err := doSomething(true)

	if fmt.Sprintf("%s", err) == ERROR_MSG_01 {
		fmt.Println("ERROR: ERROR_MSG_01に応じた処理を行う")
		os.Exit(1)

	} else if fmt.Sprintf("%s", err) == ERROR_MSG_02 {
		fmt.Println("ERROR: ERROR_MSG_02に応じた処理を行う")
		os.Exit(1)
	}
	fmt.Println("main is success")
}

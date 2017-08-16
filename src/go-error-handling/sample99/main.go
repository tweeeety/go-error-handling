package main

import "fmt"

const (
	ERROR_MSG_01 = "doSomething is error. b is true"
	ERROR_MSG_02 = "doSomething is error. b is false"
)

type MyErrorInterface interface {
	error
	IsOk() bool
}

type MyError struct {
	Msg    string
	Code   int
	Status bool
}

func (err *MyError) Error() string {
	return fmt.Sprintf("ERROR: %d %s", err.Code, err.Msg)
}

func (err *MyError) IsOk() bool {
	return err.Status
}

var (
	MyError_01 = &MyError{Msg: "this is MyError_001", Code: 30001}
	MyError_02 = &MyError{Msg: "this is MyError_002", Code: 30002}
)

type MyError2 struct {
	Msg  string
	Code int
}

func (err *MyError2) Error() string {
	return fmt.Sprintf("ERROR2: %d %s", err.Code, err.Msg)
}

func doSomething(b bool) error {

	if b {
		return MyError_01

	} else {
		MyError_02.Status = true
		return MyError_02
	}

	return nil
}

func main() {

	err := doSomething(false)

	fmt.Printf("A: %+v\n", err)
	//fmt.Printf("B: %+v, ", err.(type))
	//fmt.Printf("B: %+v\n", err.(*MyError))
	//fmt.Printf("C: %+v\n", err.(*MyError2))

	if serr, ok := err.(*MyError); ok {
		fmt.Printf("ok1: %+v\n", serr)
	}

	if serr, ok := err.(*MyError2); ok {
		fmt.Printf("ok2: %+v\n", serr)
	}
	/*
		if err == MyError_01 {
			fmt.Println("ERROR: インスタンスMyError_01に応じた処理を行う")
			os.Exit(1)

		} else if err == MyError_02 {
			fmt.Println("ERROR: インスタンスMyError_02に応じた処理を行う")

			if MyError_02.IsOk() {
				fmt.Println("ERROR: さらにStatus trueに応じた処理を行う")

			} else {
				fmt.Println("ERROR: さらにStatus falseに応じた処理を行う")
				os.Exit(1)
			}
		}
	*/

	fmt.Println("main is success")
}

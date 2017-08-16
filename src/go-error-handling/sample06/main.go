package main

import "fmt"

type MyError_01 struct {
	Code int
}

func (err *MyError_01) Error() string {
	return fmt.Sprintf("this is MyError_01")
}

type MyError_02 struct {
	Code int
}

func (err *MyError_02) Error() string {
	return fmt.Sprintf("this is MyError_02")
}

func doSomething(b bool) error {

	if b {
		return &MyError_01{Code: 30001}

	} else {
		return &MyError_02{Code: 30002}
	}

	return nil
}

func main() {

	err := doSomething(false)

	switch e := err.(type) {
	case *MyError_01:
		fmt.Println("ERROR: MyError_01の型に応じた処理")
		fmt.Printf("ERROR: err=%+v, e=%+v, code=%d\n", err, e, e.Code)

	case *MyError_02:
		//fmt.Printf("ERROR: MyError_02のエラー, err=%+v, code=%d\n", e, e.Code)
		fmt.Println("ERROR: MyError_02の型に応じた処理")
		fmt.Printf("ERROR: err=%+v, e=%+v, code=%d\n", err, e, e.Code)
	}

	fmt.Println("main is success")
}

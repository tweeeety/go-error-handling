package main

import "fmt"

const (
	ERROR_MSG_01 = "doSomething is error. b is true"
	ERROR_MSG_02 = "doSomething is error. b is false"
)

type MyError struct {
	Msg  string
	Code int
}

func (err *MyError) Error() string {
	return fmt.Sprintf("%s, %d", err.Msg, err.Code)
}

var (
	MyError_01 = &MyError{Msg: "MyError_001 is occur", Code: 30001}
	MyError_02 = &MyError{Msg: "MyError_002 is occur", Code: 30002}
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

	switch err {
	case MyError_01:
		fmt.Println("ERROR: インスタンスMyError_01に応じた処理")
		fmt.Printf("ERROR: %+v\n", err)

	case MyError_02:
		fmt.Println("ERROR: インスタンスMyError_02に応じた処理")
		fmt.Printf("ERROR: %+v\n", err)

		// ちなみにdoSomethingの返り値は
		// errorインターフェースであってMyError_01ではない。
		// ココに以下のようなコードは書けない
		//fmt.Printf("ERROR: Msg=%s, Code=%d", err.Msg, err.Code)
	}
	fmt.Println("main is success")
}

package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("doSomething is error.")
}

// nilと比較してerrorオブジェクトが返ってたらエラーと見なす
func doHoge() {
	err := doSomething()

	if err != nil {
		fmt.Println("doHoge is failed")
		return
	}

	// 処理を続ける
	fmt.Println("doHoge is success")
}

//
func doFuga() {

	_ = doSomething()

	// 処理を続ける
	fmt.Println("doFuga is success")
}

func doPiyo() {
	doSomething()

	// 処理を続ける
	fmt.Println("doPiyo is success")
}

func main() {

	doHoge()

	doFuga()

	doPiyo()

	fmt.Println("main is success")
}

package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	return errors.New("doSomething is error.")
}

func doHoge() {
	err := doSomething()

	// nilと比較してerrorオブジェクトが返ってたらエラーと見なす
	if err != nil {
		fmt.Println("doHoge is failed")
		return
	}

	// 処理を続ける
	fmt.Println("doHoge is success")
}

func doFuga() {

	// _ でerrorを無視する
	_ = doSomething()

	// 処理を続ける
	fmt.Println("doFuga is success")
}

func doPiyo() {
	// 単に受け取らないことも
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

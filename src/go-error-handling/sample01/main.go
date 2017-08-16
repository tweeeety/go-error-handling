package main

import (
	"errors"
	"fmt"
	"os"
)

func doSomething() error {
	return errors.New("doSomething is error.")
}

func main() {

	err := doSomething()

	// nilと比較してerrorオブジェクトが返ってたらエラーと見なす
	if err != nil {
		fmt.Println("main is failed")
		os.Exit(1)
	}

	fmt.Println("main is success")
}

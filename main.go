package main

import (
	"demo/injector"
	"fmt"
)

func main() {
	ret, err := injector.InitializeIjk()
	if err != nil {
		fmt.Errorf("is nil")
		return
	}
	fmt.Printf("%v\n", ret)
}

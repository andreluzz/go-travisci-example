package main

import (
	"fmt"
	"github.com/andreluzz/go-travisci-example/app"
)

func main() {
	for i:=0; i <= 15; i++ {
		fmt.Println(app.FizzBuzz(i))
	}
}
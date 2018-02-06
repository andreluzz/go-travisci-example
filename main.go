package main

import (
	"fmt"
	"github.com/andreluzz/go-travisci-example/app"
)

var version = "Development"

func main() {
	for i := 0; i <= 15; i++ {
		fmt.Println(app.FizzBuzz(i))
	}
	fmt.Println(app.ReadXml())
	fmt.Printf("Version: %s", version)
}

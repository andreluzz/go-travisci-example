package app

import (
	"strconv"
	"github.com/andreluzz/go-travisci-example/xml"
)

func FizzBuzz(num int) string {
	div3 := (num % 3) == 0
	div5 := (num % 5) == 0

	if div3 && div5 {
		return "FizzBuzz"
	} else if div3 {
		return "Fizz"
	} else if div5 {
		return "Buzz"
	} else {
		return strconv.Itoa(num)
	}
}


func ReadXml() string {
	result, _ := xml.Parser()
	return result
}
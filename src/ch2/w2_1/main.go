package main

import (
	"fmt"

	"github.com/james604s/learn-go-programming/src/ch2/w2_1/tempconv"
)

// run command:
// go run ./src/chapter2/work2_1/work2_1.go
func main() {
	var k tempconv.Kelvin
	fmt.Printf("Kelvin: %s\tCelsius: %s\t Fahrenheit: %s\n", k.String(), tempconv.KToC(k).String(), tempconv.KToF(k).String())
}

package main

import (
	"fmt"
	"os"
)

func run() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(pwd)
}

func main() {
	run()
}

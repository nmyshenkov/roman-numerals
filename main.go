package main

import (
	"flag"
	"fmt"
)

func main() {
	digit := flag.Int("d", 0, "Write digit")
	flag.Parse()

	if digit == nil {
		fmt.Println("Please, write digit")
		return
	}
	numMap := make(map[int]string)
	initNumMap(numMap)

}

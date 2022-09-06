package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return src, errors.New("missing string")
	}
	return src, nil
}

func main() {
	str, err := readInput()
	if err != nil {
		fmt.Println("0")
		return
	}
	words := strings.Split(str, " ")
	fmt.Println(len(words))
}

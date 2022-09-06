package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
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
		fail(err)
	}
	words := strings.Split(str, " ")
	fmt.Println(len(words))
}

func fail(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
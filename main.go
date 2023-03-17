package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var str, subStr string
	flag.StringVar(&str, "str", "None", "Set str")
	flag.StringVar(&subStr, "substr", "None", "Set substr")
	flag.Parse()
	fmt.Println(strings.Contains(str, subStr))
}

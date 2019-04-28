package main

import (
	"flag"
	"fmt"

	"github.com/wangyuntao/term"
)

func main() {
	flag.Parse()

	dl := newDl()
	err := term.Do(dl.loop)
	if err != nil {
		panic(err)
	}
	fmt.Println(dl.s)
}

package main

import (
	"fmt"
	"github.com/timpellison/flaggy/client"
	"os"
)

func main() {
	url := os.Args[1]
	flag := os.Args[2]
	cl := client.NewFlaggyClient(url)
	f, e := cl.GetFlag(flag)
	if e != nil {
		panic(e)
	}
	fmt.Println(f)
}

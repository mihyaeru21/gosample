package main

import (
	"fmt"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("box", "./files")
	fmt.Println("hogeee")
	s, err := box.FindString("a/b.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

package main

import (
	"fmt"
	"gitlab.com/tmlib/wana"
)

func main() {
	x := wana.GetXArray(0.01, 10, true)
	fmt.Println(x)
}

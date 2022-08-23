package main

import (
	"fmt"

	"github.com/Slimyi/FemboyDesktopFlooder/r34dl"
)

func Test() {
	str := r34dl.Fetch("femboy", 50, 2)
	fmt.Println(str)
}

package main

import (
	"fmt"
	"singlton/pkg/singl"
)

func main() {
	fmt.Println(singl.GetInstance())

}

package main

import (
	"Profile/pkg/profile"
	"fmt"
)

func main() {
	newProfile := profile.CreateProfile("user", "123", true)
	fmt.Println(newProfile)
}

package main

import (
	"fmt"

	"github.com/Trinergy/gofun/anagram"
)

func main() {
	fmt.Println(anagram.Verify("string", "trings"))
	fmt.Println(anagram.Verify("Dirty     room", "Dormitory"))
}

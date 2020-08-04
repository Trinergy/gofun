package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// IsAnagram compares two strings to determine anagramninity
func IsAnagram(x, y string) bool {
	formattedString1 := strings.ReplaceAll(strings.TrimSpace(strings.ToLower(x)), " ", "")
	formattedString2 := strings.ReplaceAll(strings.TrimSpace(strings.ToLower(y)), " ", "")

	strSlice1 := strings.Split(formattedString1, "")
	strSlice2 := strings.Split(formattedString2, "")

	sort.Strings(strSlice1)
	sort.Strings(strSlice2)

	return reflect.DeepEqual(strSlice1, strSlice2)
}

func main() {
	fmt.Println(IsAnagram("string", "trings"))
	fmt.Println(IsAnagram("Dirty     room", "Dormitory"))
}

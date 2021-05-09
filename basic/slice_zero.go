package main

import "fmt"

func main() {
	// ตัวแปร langs เป็น slice เก็บค่าเป็น string
	// ภายใน slice มีค่าเป็น {}
	langs := []string{}
	fmt.Printf("langs: %#v\n", langs)
	// check slice
	if langs == nil {
		fmt.Printf(`Yes nil "langs" is a nil slice`)
	} else {
		fmt.Printf("langs is Not nil, value: %#v\n", langs)
	}
}

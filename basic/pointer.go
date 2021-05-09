package main

import "fmt"

func main() {
	me := "Golang"
	fmt.Printf("You are %s\n", me)

	// addr เก็บที่อยู่ของตัวแปร me
	// addr เก็บค่า pointer ของ string
	var addr *string = &me
	fmt.Print(addr)
	// check type
	fmt.Printf("%T\n", addr)

	// ใส่ค่า "Penguin" ลงใน me
	// *addr จะชี้ไปยังที่อยู่ของ me
	*addr = "Penguin"
	fmt.Printf("You are %s\n", me)

}

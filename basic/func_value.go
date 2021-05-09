package main

import "fmt"

var name string = "Oil"

func say(n string) {
	fmt.Printf("My name is %s\n", n)
}
func call(op func(int, int) int) {
	r := op(4, 5)
	fmt.Printf("result : %v\n", r)
}
func main() {

	var name string = "Oil"
	fmt.Printf("value : %v\n", name)
	// func ใน go เปรียบเสมือน value สามารถโยนไปที่ไหนก็ได้
	plus := func(a int, b int) int { return a + b }

	p := plus(1, 2)
	fmt.Println(p)
	fmt.Printf("type: %T\n", plus)
	// func say เรียกใช้ตัวแปร name ที่ประกาศไว้ด้านบน
	say(name)
	// func call เรียกใช้ตัวแปร plus ได้เพราะมีโครงสร้างที่คล้ายกันคือ func(int, int) int)
	call(plus)

	minus := func(a, b int) int {
		return a - b
	}
	// func call เรียกใช้ตัวแปร minusได้เพราะมีโครงสร้างที่คล้ายกันคือ func(int, int) int)
	call(minus)

}

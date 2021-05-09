package main

import "fmt"

func main() {
	// map เหมือนกับ dictionary มี key และ value
	// ตัวแปร status เป็น map มี key เป็น int มี value เป็น string
	status := map[int]string{
		200: "OK",
		300: "Permanent Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
		66:  "test",
	}
	fmt.Printf("% #v\n", status)

	// นับขนาดใน map
	l := len(status)
	fmt.Printf("length: %d\n\n", l)
	// เปลี่ยนค่าใน map
	status[200] = "Ok"
	fmt.Printf("% #v\n", status)

	// การเพิ่มค่า ถ้ามีอยู่แล้วจะเปลี่ยนค่าเดิม แต่ถ้าไม่มีจะเพิ่มเข้าไปใหม่
	status[205] = "Test"
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	// เรียกดู map
	value := status[200]
	fmt.Printf("% #v\n", value)

	// ลบ map ในคำสั่ง delete()
	delete(status, 205)
	fmt.Printf("% #v\n", status)
	fmt.Printf("length: %d\n\n", len(status))

	// check ค่าใน map
	if v, ok := status[66]; ok {
		fmt.Printf("value : %q\n\n", v)
	} else {
		fmt.Println("not found")
	}

	// ค่าใน map คือ empty
	// ผลลัพธ์คือ m is not nil. value : map[string]string{}
	var n map[string]string = map[string]string{}
	if n == nil {
		fmt.Printf("m is nil,value : %#v\n", n)
	} else {
		fmt.Printf("m is not nil. value : %#v\n", n)
	}

	// การใช้ make
	//var m map[string]string = make(map[string]string)
	var m = make(map[string]string)
	if m == nil {
		fmt.Printf("m is nil,value : %#v\n", m)
	} else {
		fmt.Printf("m is not nil. value : %#v\n", m)
	}
}

package main

import "fmt"

func testLoop() {
	// go ไม่มี while loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// ลองเขียน for แบบ while loop
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
}
func main() {
	testLoop()
	// for ใช้วนลูปได้ทั้ง slice และ array
	langs := []string{"golang", "python", "java", "javascript"}
	fmt.Println("\nfor basic")

	// วนลูป slice
	// วนตามจำนวน slice โดยค่าเริ่มแรกเป็น 0 แล้วเพิ่มทีละ 1
	for i := 0; i < len(langs); i++ {
		// ตัวแปร value เก็บค่า slice ในแต่ละลำดับ
		value := langs[i]
		// แสดงแต่ละลำดับใน slice : ข้อมูลของตำแหน่งนั้น เช่น 0 : golang
		fmt.Println(i, ":", value)
	}
	fmt.Println("\nfor range slice")
	// range จะดึงค่ามาทีละค่าพร้อมกับค่าลำดับของข้อมูลด้วย
	// จะ return 2 ค่า คือ index,value
	for index, value := range langs {
		fmt.Println(index, ":", value)
	}
	// เอาแต่ value
	fmt.Println("\nonly value")
	for _, value := range langs {
		fmt.Println("only value : ", value)
	}
	// เอาแต่ index
	fmt.Println("\nonly index")
	for index := range langs {
		fmt.Println("only index : ", index)
	}
}

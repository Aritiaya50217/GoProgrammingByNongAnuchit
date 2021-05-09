package main

import "fmt"

func showAll(ms [4]string) {
	fmt.Printf("show all : %v\n", ms)
}

func main() {
	// array ต้องกำหนดขนาดเสมอ
	// เก็บข้อมูล 4 ตัว เป็น string
	var langs = [4]string{"golang", "python", "java", "javascript"}
	fmt.Printf("langs: %v\n", langs)

	// เรียกข้อมูลที่ต้องการ
	n := langs[1]
	fmt.Printf("%s\n", n)

	// เปลี่ยนค่าใน array
	// ชื่อตัวแปร ตามด้วย subscript แล้วใส่ตำแหน่งของข้อมูลที่ต้องการเปลี่ยน
	langs[3] = "c++"
	fmt.Printf("%v\n", langs)

	// func showAll เรียกใช้ตัวแปร langs ได้เพราะมีคุณสมบัติคล้ายกันคือ [4]string ถือว่าเป็น type เดียวกัน
	showAll(langs)

	cords := [4]string{"A", "B", "C"}
	showAll(cords)
}

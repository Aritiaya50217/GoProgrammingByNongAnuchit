package main

import "fmt"

func printSlice(ns []string) {
	fmt.Printf("printSlice : ns: %#v\n", ns)
}

func main() {

	langs := []string{"golang", "python", "java", "javascript"}
	// ใช้ # จะแสดงข้อมูลแบบมี ""
	fmt.Printf("langs: %#v\n", langs)

	// ค่าตำแหน่งที่ 0 ถึง 2 แต่ไม่เอา 2
	a := langs[0:2]
	fmt.Printf("a : %#v\n", a)
	fmt.Printf("langs : %#v\n", langs)

	b := langs[1:3]
	fmt.Printf("b : %#v\n", b)
	fmt.Printf("langs : %#v\n", langs)

	// เอาค่าแรกจนถึงตัวสุดท้าย
	c := langs[0:]
	fmt.Printf("c : %#v\n", c)
	fmt.Printf("langs : %#v\n", langs)

	// ใช้ตัวแปร langs ได้เพราะ langs เป็น slice
	printSlice(langs)
	// รับเฉพาะ type เดียวกัน ไม่สามารถรับตัวแปรที่เป็น array ได้
	// cords := [4] string{"A","B","C","D"}
	// เปลี่ยน array ให้เป็น slice เพื่อให้รับค่าได้
	cords := [4]string{"A", "B", "C", "D"}
	printSlice(cords[:])

}

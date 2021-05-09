package main

import "fmt"

func main() {
	// slice ไม่ต้องระบุขนาดใน subscript
	// slice มีขนาดที่สามารถยืดและหดได้ การหาขนาดจะใช้คำสั่ง len
	langs := []string{"golang", "python", "java", "javascript"}
	fmt.Printf("langs: %s\n", langs)

	// เรียกดูค่าแต่ละตำแหน่งใน slice
	n := langs[0]
	fmt.Printf("langs[0] : %s\n", n)

	// เปลี่ยนแปลงค่าในแต่ละตำแหน่ง
	langs[3] = "c++"
	fmt.Printf("langs: %v\n", langs)

	// หาขนาด
	l := len(langs)
	fmt.Println("length of langs :", l)

	// เพิ่มขนาดของ slice ใช้คำสั่ง append
	// append ไม่มีการเปลี่ยนแปลง slice ต้นทางแต่จะ return ค่าใหม่เพิ่มเข้าไป
	langs = append(langs, "F#")
	fmt.Printf("langs : %v\n", langs)
	// หาขนาด
	fmt.Printf("length of langs : %v\n", len(langs))

}

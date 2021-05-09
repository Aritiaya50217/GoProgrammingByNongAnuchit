package main

import (
	"fmt"
	"log"
)

// ถ้า func นั้นมี effact ค่าพารามิเตอร์ตัวสุดท้ายจะเป็น error
func ReadFile(name string) (string, error) {
	// error ใน go เป็น type interface เช่น var err error ตัว error คือ type
	// error ใน Go คือ string เฉยๆ
	// สร้าง error ใหม่ใช้ package errors.New
	// var err error = errors.New("file not found !!")
	// fmt.Println(err)
	return "data... ", nil

}

func main() {
	// อ่านค่าไฟล์ profile.txt
	data, err := ReadFile("profile.txt")
	// err เป็น type interface ค่า zerovalue ของ interface คือ nil
	if err != nil {
		log.Println(err)
	}
	fmt.Println("read file success :", data)
}

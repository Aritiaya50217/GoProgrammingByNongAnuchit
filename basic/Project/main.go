package main

import (
	"fmt"
	"github.com/Artit50217/igapp/user"
	"github.com/Artit50217/igapp/time"
)

/* ย้ายไป package user
func info(name string, msg string, age int) {
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My message is %s\n", msg)
	fmt.Printf("I'm %d year old\n", age)
}
*/

/* ย้ายไป package time
func today() string {
	return "Hello"
}
*/

func main() {
	// user คือ packageใหม่ที่เราสร้างขึ้น
	// ใช้เมื่อต้องการดึง package จากโฟล์เดอร์อื่นมา
	user.Info("Oil", "Golang", 23)
	// time คือ packageใหม่ที่เราสร้างขึ้น
	t := time.Today()
	fmt.Println("today is", t)
}

package main

import "fmt"

func info(name string, msg string) {
	fmt.Printf("My name is %s\n", name)
}
func infoo(name, msg string, age int) {
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My message is %s\n", msg)
	fmt.Printf("I am %d year old\n", age)
}
func today() string {
	return "Hi"
}

// คืนค่ามากกว่า 1 ค่าให้ใส่ ()
func swap(x, y string) (string, string) {
	return x, y
}
func main() {
	info("Golang", "!!!")
	info("Oil", "...")
	infoo("Oil", "Hello", 23)

	day := today()
	fmt.Println(day)

	// สร้างตัวแปรรับค่าพารรามิเตอร์จาก function
	a, b := swap("Artitaya", "Yaemjaraen")
	fmt.Println(a, b)

}

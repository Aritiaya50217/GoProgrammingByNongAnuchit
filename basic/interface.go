package main

import (
	"fmt"
)

// interface ข้อตกลงร่วมกัน ว่าต้องมีลักษณะแบบนี้เท่านั้น
// interface ใน Go เป็น m:n
type Phone interface {
	Call(number string)
}
type Samsung struct {
	Name string
}

// func มี Call เป็น Method ของ Sumsung
func (s Samsung) Call(number string) {
	fmt.Println(s.Name, "calling", number)
}

// func มี Answer เป็น Method ของ Sumsung
func (s Samsung) Answer() {
	fmt.Println(s.Name, "Hello there !!")
}
func Dial(p Phone) {
	p.Call("+6678787888")
}
func main() {
	s := Samsung{
		Name: "S10",
	}
	Dial(s)
	// call func แบบ method
	s.Answer()
}

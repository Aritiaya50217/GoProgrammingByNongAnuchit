package main

import "fmt"

/*
	The zero value is :
	0 for number types,
	false for the boolean type
	"" (the empty string) for string
*/
/* basic type
bool
string
int int8 int16 int32 int64
unit uint8 uint16 uint32 uint64 unitptr

byte // alias for unit8
rune คล้ายกับ charactor คือ เป็น 1 ตัวอักษร
// alias for int32
// represnts a Unicode code point
float32 float64
complex64 complex128

*/
func main() {
	var name string = "Golang"
	v := 55
	fmt.Printf("name: %v\n", name)
	// %T ประภทของตัวแปร
	fmt.Printf("type: %T\n", name)
	fmt.Print(v)
}

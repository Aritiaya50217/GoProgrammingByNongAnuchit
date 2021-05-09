package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Monday
	switch today {
	case time.Monday:
		fmt.Println("today is Monday")
		// fallthrough ใช้เมื่อต้องการให้ทำ case ข้างล่างด้วย
		fallthrough
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Printf("สวัสดี %v\n", today)
	}
}

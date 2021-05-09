package main

import "fmt"

func main() {
	// const คือ การประกาศค่าตัวแปรที่ไม่สามารถ เปลี่ยนแปลงค่าได้อีก
	const day = "Monday"
	const (
		One   = 1
		Two   = 2
		Three = 3
	)
	const (
		// iota จะ auto assign ค่าต่อไปโดยเพิ่มทีละ 1 ค่าแรกจะเป็น 0
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println("day :", day)
	fmt.Println("One :", One)
	fmt.Println("Two :", Two)
	fmt.Println("Three :", Three)
	fmt.Println("Sunday :", Sunday)
	fmt.Println("Monday :", Monday)
	fmt.Println("Tuesday :", Tuesday)
	fmt.Println("Wednesday :", Wednesday)
	fmt.Println("Thursday :", Thursday)
	fmt.Println("Friday :", Friday)
	fmt.Println("Saturday :", Saturday)

}

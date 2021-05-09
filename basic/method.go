package main

import "fmt"

type User struct {
	Username      string
	FullName      string
	ProfilePicUrl string
}

func info(u User) {

	fmt.Printf("User name : %#v\n", u.Username)
	fmt.Printf("Full nmae : %#v\n", u.FullName)
	fmt.Printf("Profile Picture URL : %#v\n", u.ProfilePicUrl)

}

// เปลี่ยน func info เป็น Method ของ type User
// Info ตัวแรกเป็นพิมพ์ใหญ่ คือ Public
// info ตัวแรกเป็นพิมพ์เล็ก คือ Private
func (u User) info() {

	fmt.Printf("User name : %#v\n", u.Username)
	fmt.Printf("Full nmae : %#v\n", u.FullName)
	fmt.Printf("Profile Picture URL : %#v\n", u.ProfilePicUrl)

}
func main() {
	// ประกาศตัวแปร u มี type เป็น User
	u := User{
		Username:      "golang",
		FullName:      "Basic Golang",
		ProfilePicUrl: "https://knowhere.fake/gopher.jpg",
	}
	// เรียกใช้ func
	info(u)
	// เรียกใช้ func ที่เป็น method
	u.info()
}

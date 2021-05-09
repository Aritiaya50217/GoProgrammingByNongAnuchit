package main

import "fmt"

type User struct {
	Username    string
	FullName    string
	ProfileName string
}

func main() {
	username := "golang"
	fullName := "Basic Golang"
	profilePicUrl := "https://knowhere.fake/gopher.jpg"
	fmt.Println(username, fullName, profilePicUrl)

	/* เรียกใช้ struct แบบที่1
	u := User{}
	// u ชี้ไปที่ field แต่ละตัวใน struct ที่เราสร้าง
	u.Username = username
	u.FullName = fullName
	u.ProfileName = profilePicUrl
	fmt.Printf("%#v\n", u)
	*/
	u := User{
		Username:    username,
		FullName:    fullName,
		ProfileName: profilePicUrl,
	}
	fmt.Printf("%#v\n", u)
	// ดึงค่าจาก struct ที่เราต้องการ
	name := u.FullName
	fmt.Println(name)
	// แบบ inline
	fmt.Println(u.FullName)
}

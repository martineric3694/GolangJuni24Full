package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Structs untuk tabel
type User struct {
	gorm.Model
	Name      string
	Profile   Profile
	ProfileID uint
	Posts     []Post
	Roles     []Role `gorm:"many2many:user_roles"`
}

type Profile struct {
	gorm.Model
	Phone  string
	UserID uint
}

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
}

type Role struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_roles"`
}

func main() {
	// Koneksi ke database MySQL
	dsn := "root:root@tcp(127.0.0.1:3306)/JUN24?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrasi tabel
	db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Role{})

	// Menambahkan data contoh
	user := User{
		Name: "John Doe",
		Profile: Profile{
			Phone: "123-456-7890",
		},
		Posts: []Post{
			{Title: "First Post", Content: "Content of the first post"},
			{Title: "Second Post", Content: "Content of the second post"},
		},
		Roles: []Role{
			{Name: "Admin"},
			{Name: "User"},
		},
	}
	db.Create(&user)

	// Membaca data user beserta relasi
	var readUser User
	db.Preload("Profile").Preload("Posts").Preload("Roles").First(&readUser, user.ID)
	fmt.Printf("User: %v\n", readUser)
	fmt.Printf("Profile: %v\n", readUser.Profile)
	fmt.Printf("Posts: %v\n", readUser.Posts)
	fmt.Printf("Roles: %v\n", readUser.Roles)

	db.Delete(readUser).Where(1)

}

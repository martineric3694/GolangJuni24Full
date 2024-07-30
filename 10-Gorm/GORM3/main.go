package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Structs untuk tabel
type User struct {
	gorm.Model
	Name    string
	Profile Profile
	Posts   []Post
}

type Profile struct {
	gorm.Model
	UserID  uint
	Email   string
	Address string
}

type Post struct {
	gorm.Model
	UserID   uint
	Title    string
	Content  string
	Comments []Comment
	Tags     []Tag `gorm:"many2many:post_tags"`
}

type Comment struct {
	gorm.Model
	PostID  uint
	Content string
}

type Tag struct {
	gorm.Model
	Name  string
	Posts []Post `gorm:"many2many:post_tags"`
}

func main() {
	// Koneksi ke database MySQL
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrasi tabel
	db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Comment{}, &Tag{})

	// Menambahkan data contoh
	user := User{
		Name: "John Doe",
		Profile: Profile{
			Email:   "john@example.com",
			Address: "123 Main St",
		},
		Posts: []Post{
			{
				Title:   "First Post",
				Content: "This is the first post.",
				Comments: []Comment{
					{Content: "Great post!"},
					{Content: "Thanks for sharing."},
				},
				Tags: []Tag{
					{Name: "Golang"},
					{Name: "Programming"},
				},
			},
		},
	}
	db.Create(&user)

	// Membaca data User beserta Profile, Posts, Comments, dan Tags
	var readUser User
	db.Preload("Profile").Preload("Posts.Comments").Preload("Posts.Tags").First(&readUser, user.ID)
	fmt.Printf("User: %v\n", readUser)
	fmt.Printf("Profile: %v\n", readUser.Profile)
	for _, post := range readUser.Posts {
		fmt.Printf("Post: %v\n", post)
		fmt.Printf("Comments: %v\n", post.Comments)
		fmt.Printf("Tags: %v\n", post.Tags)
	}
}

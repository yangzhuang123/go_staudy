package main

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	Id        int    `db:"id" json:"id"`
	Name      string `db:"name"  gorm:"not null;index"`
	PostCount int    `db:"post_count"  gorm:"not null;index"`
}

func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户的文章数量，自增1
	return tx.Model(&User{}).Where("id = ?", p.UserId).Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	err := tx.Model(&Comment{}).Where("post_id =?", c.PostId).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		fmt.Println(c.PostId)
		return tx.Model(&Post{}).Where("id = ?", c.PostId).Update("status", "无评论").Error
	}
	return nil
}

type Post struct {
	Id       int       `db:"id" json:"id"`
	Title    string    `db:"title" json:"title" gorm:"not null"`
	Content  string    `db:"content" json:"content" gorm:"not null"`
	UserId   int       `db:"user_id" json:"user_id" gorm:"not null;index"`
	Status   string    `db:"status" json:"status"`
	User     User      `gorm:"foreignkey:UserId" json:"author"` // 添加用户关联
	Comments []Comment `gorm:"foreignkey:PostId" json:"comments"`
}

// Comment 评论
type Comment struct {
	Id      int    `db:"id" json:"id" gorm:"primary_key"`
	PostId  int    `db:"post_id" json:"postId" gorm:"not null;index"`
	Post    Post   `gorm:"foreignkey:PostId;association_foreignkey:Id"`
	UserId  int    `db:"from_user_id" json:"fromUserId" gorm:"not null;index"`
	User    User   `gorm:"foreignkey:UserId" json:"fromUser"`
	Content string `db:"content" json:"content" gorm:"not null"`
}

func RunQuery(db *gorm.DB) {
	var user User
	db.Where("id=?", 1).Find(&user)
	fmt.Println(user)

	var posts []Post
	db.Where("user_id=?", 1).Preload("User").Preload("Comments").Find(&posts)
	for _, post := range posts {
		fmt.Println(post)
	}

	var post Post
	db.Where(" id in (select post_id from (select post_id,count(1) as postCount" +
		" from comments group by post_id order by postCount desc limit 1) a)").Preload("User").
		Preload("Comments").First(&post)

	fmt.Println(post)
}

func RunHook(db *gorm.DB) {
	//db.AutoMigrate(&User{}, &Post{}, &Comment{})

	user := User{
		Name:      "<UNK>",
		PostCount: 0,
	}

	posts := []Post{
		{
			Title:   "title 1",
			Content: "<UNK>",
			UserId:  1,
			Status:  "<UNK>",
			Comments: []Comment{
				{
					PostId:  1,
					UserId:  1,
					Content: "test",
				},
				{
					PostId:  1,
					UserId:  1,
					Content: "test",
				},
			},
		},
	}

	db.Create(&user)
	db.Create(&posts)
	var comment Comment
	db.Find(&comment)
	db.Delete(&comment)

}

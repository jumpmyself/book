package model

import "time"

// Admin undefined
type Admin struct {
	ID          int64     `json:"id" gorm:"id"`
	Name        string    `json:"name" gorm:"name"`
	Password    string    `json:"password" gorm:"password"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*Admin) TableName() string {
	return "admin"
}

type Book struct {
	ID          int64     `json:"id" gorm:"id"`
	Uid         int64     `json:"uid" gorm:"uid"`
	Name        string    `json:"name" gorm:"name"`
	Cate        string    `json:"cate" gorm:"cate"`
	Status      int64     `json:"status" gorm:"status"`
	Num         int64     `json:"num" gorm:"num"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*Book) TableName() string {
	return "book"
}

// BookInfo undefined
type BookInfo struct {
	ID                 int64     `json:"id" gorm:"id"`                                   // 书的id
	BookName           string    `json:"book_name" gorm:"book_name"`                     // 书名
	Author             string    `json:"author" gorm:"author"`                           // 作者
	PublishingHouse    string    `json:"publishing_house" gorm:"publishing_house"`       // 出版社
	Translator         string    `json:"translator" gorm:"translator"`                   // 译者
	PublishDate        time.Time `json:"publish_date" gorm:"publish_date"`               // 出版时间
	Pages              int64     `json:"pages" gorm:"pages"`                             // 页数
	Isbn               string    `json:"ISBN" gorm:"ISBN"`                               // ISBN号码
	Price              float64   `json:"price" gorm:"price"`                             // 价格
	BriefIntroduction  string    `json:"brief_introduction" gorm:"brief_introduction"`   // 内容简介
	AuthorIntroduction string    `json:"author_introduction" gorm:"author_introduction"` // 作者简介
	ImgUrl             string    `json:"img_url" gorm:"img_url"`                         // 封面地址
	DelFlg             int64     `json:"del_flg" gorm:"del_flg"`                         // 删除标识
}

// TableName 表名称
func (*BookInfo) TableName() string {
	return "book_info"
}

// BookUser undefined
type BookUser struct {
	ID          int64     `json:"id" gorm:"id"`
	UserId      int64     `json:"user_id" gorm:"user_id"`
	BookId      int64     `json:"book_id" gorm:"book_id"`
	Status      int64     `json:"status" gorm:"status"`
	Time        int64     `json:"time" gorm:"time"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*BookUser) TableName() string {
	return "book_user"
}

// User undefined
type User struct {
	ID          int64     `json:"id" gorm:"id"`
	Name        string    `json:"name" gorm:"name"`
	Password    string    `json:"password" gorm:"password"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*User) TableName() string {
	return "user"
}

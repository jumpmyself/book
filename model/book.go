package model

import (
	"errors"
	"fmt"
	"time"
)

func CreateBook(b *Book) error {
	return Conn.Create(b).Error
}

func GetBook(id int64) (*Book, error) {
	var ret Book
	Conn.Where("id = ?", id).First(&ret)
	return &ret, nil
}

// GetBooksByCursor 根据起始 UID 和结束 UID 进行游标分页查询图书列表
func GetBooksByCursor(cursor, pageSize int) ([]*BookInfo, error) {
	var books []*BookInfo

	// 查询数据库
	result := Conn.Where("id > ?", cursor).Limit(pageSize).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func SaveBook(data *Book) error {
	return Conn.Save(data).Error
}

func DeleteBook(id int64) error {
	return Conn.Where("id = ?", id).Delete(nil).Error
}

func BorrowBook(uid, id int64) error {
	tx := Conn.Begin()
	//查询用户是否存在
	var user User
	tx.Where("id = ?", uid).First(&user)
	if user.ID == 0 {
		tx.Rollback()
		return errors.New("用户信息不存在")
	}

	//查询图书是否存在，是否正常
	var book Book
	tx.Where("id = ?", id).First(&book)
	if book.ID == 0 || book.Num <= 0 {
		tx.Rollback()
		return errors.New("图书信息不存在或库存不足")
	}
	//创建借阅记录
	now := time.Now()

	bu := BookUser{
		UserId:      uid,
		BookId:      id,
		Status:      1,
		Time:        0,
		CreatedTime: now,
		UpdatedTime: now,
	}
	if tx.Create(&bu).Error != nil {
		tx.Rollback()
		return errors.New("创建一个借阅记录")
	}
	//扣减图书库存
	book.Num = book.Num - 1
	if tx.Save(&book).Error != nil {
		tx.Rollback()
		return errors.New("扣减图书库存")
	}

	tx.Commit()
	return nil
}

func ReturnBook(uid, id int64) error {
	tx := Conn.Begin()
	//查询用户是否存在
	var user User
	tx.Where("id = ?", uid).First(&user)
	if user.ID == 0 {
		tx.Rollback()
		return errors.New("用户信息不存在")
	}

	//查询图书是否存在，是否正常
	var book Book
	tx.Where("id = ?", id).First(&book)
	if book.ID == 0 || book.Num <= 0 {
		tx.Rollback()
		return errors.New("图书信息不存在或库存不足")
	}

	//查询借书记录是否存在
	var bu BookUser
	tx.Where("user_id = ? and book_id = ?", uid, id).First(&bu)
	if bu.ID <= 0 {
		tx.Rollback()
		return errors.New("借阅记录不存在")
	}

	//更新借阅状态
	bu.Status = 0
	if err := tx.Save(&bu).Error; err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("修改借阅记录失败：%s", err.Error()))
	}

	//更新图书库存
	book.Num = book.Num + 1
	if err := tx.Save(&book).Error; err != nil {
		tx.Rollback()
		return errors.New(fmt.Sprintf("增加库存失败：%s", err.Error()))
	}
	tx.Commit()
	return nil
}

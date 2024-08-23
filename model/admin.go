package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func GetAdmin(name string) (Admin, error) {
	var admin Admin
	err := Conn.Where("name = ?", name).First(&admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Admin{}, fmt.Errorf("admin not found with name: %s", name)
		}
		return Admin{}, err
	}
	return admin, nil
}

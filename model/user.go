package model

// GetUser 根据用户名从数据库中获取用户信息
func GetUser(username string) (*User, error) {
	var user User
	// 执行数据库查询操作
	result := Conn.Where("name = ?", username).First(&user)
	// 检查是否发生错误
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

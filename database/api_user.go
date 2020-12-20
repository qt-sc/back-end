package  database


//GetAllUser 获取所有用户
func (dbservice *DBService) GetAllUser() ([]model.User, error) {
	var userlist []model.User
	if err := db.Table("user").Find(&userlist).Error; err != nil {
		return userlist, err
	}
	return userlist, nil
}

//GetOneUser 获取指定用户
func (dbservice *DBService) GetOneUser(name string) (model.User, error) {
	var user model.User
	if err := db.Table("user").Where("name = ?", name).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

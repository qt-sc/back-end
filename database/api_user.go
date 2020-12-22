package  database

import "github.com/qt-sc/server/model"

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

//CreateUser 创建用户
func (dbservice *DBService) CreateUser(user model.User) (bool, error) {
	
	if err := db.Table("user").Create(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

//DeleteUser 删除用户
func (dbservice *DBService) DeleteUser(user_id int64) (bool, error) {

	if err := db.Table("user").Delete(&model.User{}, user_id).Error; err != nil {
		return false, err
	}
	return true, nil
}

//UpdateUser 更新用户
func (dbservice *DBService) UpdateUser(user model.User) (bool, error) {

	// 暂定
	if err := db.Table("article").Model(&user).Updates(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

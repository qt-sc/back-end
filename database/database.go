package database

import (
	"fmt"
	"github.com/qt-sc/server/model"
)

type DBServiceInterface interface {
	GetAllPublicEssay() ([]model.Essay, error)
	UserCreateEssay(essay model.Essay) (bool, error)
	AddEssayFromZhihu(essay model.Essay) (bool, error)
}

type DBService struct {}

//GetAllPublicEssay 获取所有用户的public博客
func (dbservice *DBService) GetAllPublicEssay() ([]model.Essay, error) {
	var essaylist []model.Essay
	if err := db.Table("essay").Order("create_time").Find(&essaylist).Error; err != nil {
		return essaylist, err
	}
	for i, j := 0, len(essaylist)-1; i < j; i, j = i+1, j-1 {
		essaylist[i], essaylist[j] = essaylist[j], essaylist[i]
	}
	return essaylist, nil
}

//UserCreateEssay 用户新建博客
func (dbservice *DBService) UserCreateEssay(essay model.Essay) (bool, error) {
	fmt.Println(essay)
	if err := db.Table("essay").Create(&essay).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (dbservice *DBService) AddEssayFromZhihu(essay model.Essay) (bool, error) {
	fmt.Println(essay)
	if err := db.Table("essay").Create(&essay).Error; err != nil {
		return false, err
	}
	return true, nil
}
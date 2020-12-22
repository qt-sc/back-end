package database

import "github.com/qt-sc/server/model"

//CreateTag 创建标签
func (dbservice *DBService) CreateTag(tag model.Tag) (bool, error) {
	
	if err := db.Table("tag").Create(&tag).Error; err != nil {
		return false, err
	}

	return true, nil
}

//GetAllTag 获取所有标签
func (dbservice *DBService) GetAllTag() ([]model.Tag, error) {
	var taglist []model.Tag
	if err := db.Table("tag").Find(&taglist).Error; err != nil {
		return taglist, err
	}
	return taglist, nil
}

//GetTagByArticle 获取文章所有标签
func (dbservice *DBService) GetTagByArticle(article_id int64) ([]model.Tag, error) {
	var taglist []model.Tag

	var article model.Article
	if err := db.Table("article").Where("id = ?", article_id).First(&article).Error; err != nil {
		return taglist, err
	}

	if err := db.Model(&article).Related(&taglist, "Tags").Error; err != nil {
		return taglist, err
	}
	return taglist, nil
}

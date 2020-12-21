package  database

import "github.com/qt-sc/server/model"

//GetAllArticle 获取所有文章
func (dbservice *DBService) GetAllArticle() ([]model.Article, error) {
	var articlelist []model.Article
	if err := db.Table("article").Order("create_time desc").Find(&articlelist).Error; err != nil {
		return articlelist, err
	}
	return articlelist, nil
}

//GetArticleByUser 获取用户所有文章
func (dbservice *DBService) GetArticleByUser(user_id int64) (model.Article, error) {
	var articlelist []model.Article

	var user model.User
	if err := db.Table("user").Where("id = ?", user_id).First(&user).Error; err != nil {
		return articlelist, err
	}

	if err := db.Model(&user).Related(&articlelist).Error; err != nil {
		return articlelist, err
	}
	return articlelist, nil
}

//GetArticleByTag 获取标签所有文章
func (dbservice *DBService) GetArticleByTag(tag_name string) (model.Article, error) {
	var articlelist []model.Article

	var tag model.Tag
	if err := db.Table("tag").Where("name = ?", tag_name).First(&tag).Error; err != nil {
		return articlelist, err
	}

	//不知道能不能成功
	if err := db.Model(&tag).Related(&articlelist, "Tags").Error; err != nil {
		return articlelist, err
	}
	return articlelist, nil
}
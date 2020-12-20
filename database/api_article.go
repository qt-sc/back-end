package  database

//GetAllArticle 获取所有文章
func (dbservice *DBService) GetAllArticle() ([]model.Article, error) {
	var articlelist []model.Article
	if err := db.Table("article").Order("create_time desc").Find(&articlelist).Error; err != nil {
		return articlelist, err
	}
	return articlelist, nil
}

//GetArticleByUser 获取用户所有文章
func (dbservice *DBService) GetArticleByUser(id int64) (model.Article, error) {
	var articlelist []model.Article

	var user model.User
	if err := db.Table("user").Where("id = ?", id).First(&user).Error; err != nil {
		return articlelist, err
	}

	if err := db.Model(&user).Related(&articlelist).Error; err != nil {
		return articlelist, err
	}
	return articlelist, nil
}
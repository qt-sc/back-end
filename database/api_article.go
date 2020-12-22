package  database

import (
	"github.com/qt-sc/server/lib"
	"github.com/qt-sc/server/model"
)

//GetAllArticle 获取所有文章
func (dbservice *DBService) GetAllArticle() ([]model.Article, error) {
	var articlelist []model.Article
	if err := db.Table("article").Find(&articlelist).Error; err != nil {
		return articlelist, err
	}
	return articlelist, nil
}

//GetArticleByUser 获取用户所有文章
func (dbservice *DBService) GetArticleByUser(user_id int64) ([]model.Article, error) {
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
func (dbservice *DBService) GetArticleByTag(tag_name string) ([]model.Article, error) {
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

//GetArticleByArticle 获取指定文章
func (dbservice *DBService) GetArticleByArticle(article_id int64) (model.Article, error) {

	var article model.Article
	if err := db.Table("article").Where("id = ?", article_id).First(&article).Error; err != nil {
		return article, err
	}
	return article, nil

}

//CreateArticle 创建文章
func (dbservice *DBService) CreateArticle(article model.Article) (bool, error) {

	article.Title = lib.RemoveNonBmpUnicode(article.Title)
	article.Content = lib.RemoveNonBmpUnicode(article.Content)

	if err := db.Table("article").Create(&article).Error; err != nil {
		return false, err
	}
	return true, nil

}

//DeleteArticle 删除文章
func (dbservice *DBService) DeleteArticle(article_id int64) (bool, error) {

	if err := db.Table("article").Delete(&model.Article{}, article_id).Error; err != nil {
		return false, err
	}
	return true, nil

}

//UpadteArticleLikeNum 更新文章点赞数
func (dbservice *DBService) UpadteArticleLikeNum(article_id int64, like_num_inc int64) (bool, error) {

	var article model.Article
	if err := db.Table("article").Where("id = ?", article_id).First(&article).Error; err != nil {
		return false, err
	}

	like_num := article.LikeNum + like_num_inc;

	if err := db.Table("article").Model(&article).Update("LikeNum", like_num).Error; err != nil {
		return false, err
	}
	return true, nil

}

//UpadteArticleContent 更新文章内容
func (dbservice *DBService) UpadteArticleContent(article_id int64, content string) (bool, error) {

	var article model.Article
	if err := db.Table("article").Where("id = ?", article_id).First(&article).Error; err != nil {
		return false, err
	}

	if err := db.Table("article").Model(&article).Update("Content", content).Error; err != nil {
		return false, err
	}
	return true, nil
}
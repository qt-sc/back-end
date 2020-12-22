package database

import (
	"github.com/qt-sc/server/model"
)

type DBServiceInterface interface {
	//GetAllPublicEssay() ([]model.Article, error)
	//UserCreateEssay(article model.Article) (bool, error)
	//AddEssayFromZhihu(article model.Article) (bool, error)

	GetAllUser() ([]model.User, error)
	GetOneUser(string) (model.User, error)
	CreateUser(model.User) (bool, error)
	DeleteUser(int64) (bool, error)

	GetAllArticle() ([]model.Article, error)
	GetArticleByUser(int64) ([]model.Article, error)
	GetArticleByTag(string) ([]model.Article, error)
	GetArticleByArticle(int64) (model.Article, error)
	CreateArticle(model.Article) (bool, error)
	DeleteArticle(int64) (bool, error)
	UpdateArticleLikeNum(int64, int64) (bool, error)
	UpdateArticleContent(int64, string) (bool, error)

	GetReply(int64) (model.Reply, error)
	GetReplyByArticle(int64) ([]model.Reply, error)
	CreateReply(model.Reply) (bool, error)
	UpdateReplyLikeNum(int64, int64) (bool, error)

	GetAllTag() ([]model.Tag, error)
	GetTagByArticle(int64) ([]model.Tag, error)
}

type DBService struct {}

//GetAllPublicEssay 获取所有用户的public博客
//func (dbservice *DBService) GetAllPublicEssay() ([]model.Essay, error) {
//	var essaylist []model.Essay
//	if err := db.Table("essay").Order("create_time").Find(&essaylist).Error; err != nil {
//		return essaylist, err
//	}
//	for i, j := 0, len(essaylist)-1; i < j; i, j = i+1, j-1 {
//		essaylist[i], essaylist[j] = essaylist[j], essaylist[i]
//	}
//	return essaylist, nil
//}
//
////UserCreateEssay 用户新建博客
//func (dbservice *DBService) UserCreateEssay(essay model.Essay) (bool, error) {
//	fmt.Println(essay)
//	if err := db.Table("essay").Create(&essay).Error; err != nil {
//		return false, err
//	}
//	return true, nil
//}
//
//func (dbservice *DBService) AddEssayFromZhihu(essay model.Essay) (bool, error) {
//	fmt.Println(essay)
//	if err := db.Table("essay").Create(&essay).Error; err != nil {
//		return false, err
//	}
//	return true, nil
//}
package database

import "github.com/qt-sc/server/model"

//CreateReply 创建评论
func (dbservice *DBService) CreateReply(reply model.Reply) (bool, error) {
	
	if err := db.Table("reply").Create(&reply).Error; err != nil {
		return false, err
	}

	return true, nil
}

//GetReply 获取指定评论
func (dbservice *DBService) GetReply(reply_id int64) (model.Reply, error) {
	var reply model.Reply
	if err := db.Table("reply").Where("id = ?", reply_id).First(&reply).Error; err != nil {
		return reply, err
	}
	return reply, nil
}

//GetReplyByArticle 获取文章所有评论
func (dbservice *DBService) GetReplyByArticle(article_id int64) ([]model.Reply, error) {
	var replylist []model.Reply

	var article model.Article
	if err := db.Table("article").Where("id = ?", article_id).First(&article).Error; err != nil {
		return replylist, err
	}

	if err := db.Model(&article).Related(&replylist, "Replies").Error; err != nil {
		return replylist, err
	}
	return replylist, nil
}

//UpadteReplyLikeNum 更新文章点赞数
func (dbservice *DBService) UpadteReplyLikeNum(reply_id int64, like_num_inc int64) (bool, error) {

	var reply model.Reply
	if err := db.Table("reply").Where("id = ?", reply_id).First(&reply).Error; err != nil {
		return false, err
	}

	like_num := reply.LikeNum + like_num_inc;

	if err := db.Table("reply").Model(&reply).Update("LikeNum", like_num).Error; err != nil {
		return false, err
	}
	return true, nil

}

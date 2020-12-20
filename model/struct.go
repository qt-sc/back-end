package model

//User 定义在数据库中保存的user数据结构
type User struct {
	//Username string `gorm:"username"`
	//Password string `gorm:"password"`
	//Email    string `gorm:"email;PRIMARY_KEY"`
}

//Essay 保存发布内容信息
type Essay struct {
	ID         string `json:"id" gorm:"id;PRIMARY_KEY"`
	CreatorUrl string `json:"creatorurl" gorm:"creator_url"`
	CreateTime string `json:"createtime" gorm:"create_time"`
	//IsPublic     string   `json:"ispublic" gorm:"is_public"`
	Content string `json:"content" gorm:"content"`
	//PictureName  string `json:"picturename" gorm:"picture_name"`
	GoodCount int `json:"goodcount" gorm:"good_count"`
	ReadCount int `json:"readcount" gorm:"read_count"`
}

//CommentItem 保存评论
type CommentItem struct {
	ID               string `json:"id" gorm:"id;PRIMARY_KEY"`
	FromUserEmail    string `json:"from_user_email" gorm:"from_user_email"`
	TargetID         string `json:"target_blog_id" gorm:"target_id"`
	Content          string `json:"content" gorm"content"`
	TargetBlogID     string `json:"at_blog_id" gorm:"target_blog_id"`
	TargetCommentcID string `json:"target_commentc_id" gorm:"target_commentc_id"`
}

package model

type Article struct {

	Id int64 `json:"id,omitempty" gorm:"id"`

	Title string `json:"title,omitempty" gorm:"title"`

	ReadNum int64 `json:"readNum,omitempty" gorm:"read_num"`
 
	LikeNum int64 `json:"likeNum,omitempty" gorm:"like_num"`
 
	Content string `json:"content,omitempty" gorm:"type:text;content"`

	UserID int64 `json:"user_id,omitempty" gorm:"user_id"`

	Replies []Reply `json:"replies,omitempty" gorm:"replies"`
 
	Tags []Tag `json:"tags,omitempty" gorm:"tags;many2many:article_tags`

	Url string `json:"url,omitempty" gorm:"url"`
}
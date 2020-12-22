package model

import (
	"time"
)
 
type Reply struct {

	Id int64 `json:"id,omitempty" gorm:"id"`

	ArticleID int64 `json:"article_id,omitempty" gorm:"article_id"`

	LikeNum int64 `json:"likeNum,omitempty" gorm:"like_num"`
 
	CreateTime time.Time `json:"createTime,omitempty" gorm:"create_time"`
 
	Content string `json:"content,omitempty" gorm:"content"`
 
	AuthorUrl string `json:"author-url,omitempty" gorm:"author_url"`

	Url string `json:"url,omitempty" gorm:"url"`
}
 
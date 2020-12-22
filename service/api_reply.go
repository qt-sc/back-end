package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/qt-sc/server/lib"
	"github.com/qt-sc/server/model"
)

func CreateReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	r.ParseForm()
	idstr := r.PostFormValue("id")
	reply_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}

	reply_article_id, err := strconv.Atoi(r.PostFormValue("article_id"))
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}

	reply_likeNum, err :=  strconv.Atoi(r.PostFormValue("likeNum"))
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}

	reply_createTime, err := time.ParseInLocation("2006-01-02 15:04:05", r.PostFormValue("createTime"), time.Local)
	if err != nil {
		log.Fatal("string to time fail", err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	reply := model.Reply{
		Id:       	int64(reply_id),
		ArticleID:	int64(reply_article_id),
		LikeNum: 	int64(reply_likeNum),
		CreateTime: reply_createTime,
		Content:    r.PostFormValue("content"),
		AuthorUrl:	r.PostFormValue("author_url"),
		Url:      	r.PostFormValue("url"),
	}

	_, err = dbServer.CreateReply(reply)
	if err != nil {
		log.Fatal("Fail to create reply", err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	w.WriteHeader(http.StatusOK)
}

func GetReplies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "articles")
	article_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}
	replies, err := dbServer.GetReplyByArticle(int64(article_id))
	if err != nil {
		log.Fatal("Fail to get replies by article ID", err)
		w.WriteHeader(http.StatusNotFound)
	}

	msg, err := json.Marshal(replies)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)

	w.WriteHeader(http.StatusOK)
}

func GetReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "replies")
	reply_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}
	reply, err := dbServer.GetReply(int64(reply_id))
	if err != nil {
		log.Fatal("Fail to get reply by ID", err)
		w.WriteHeader(http.StatusNotFound)
	}

	msg, err := json.Marshal(reply)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)

	w.WriteHeader(http.StatusOK)
}

func LikeReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "replies")
	reply_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_, err = dbServer.UpdateReplyLikeNum(int64(reply_id), int64(1))
	if err != nil {
		log.Fatal("Fail to update reply's like number by ID", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

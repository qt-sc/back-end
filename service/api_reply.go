package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/qt-sc/server/lib"
	"github.com/qt-sc/server/model"
)

func CreateReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	body, _ := ioutil.ReadAll(r.Body)

	var reply model.Reply
	json.Unmarshal([]byte(body), &reply)

	_, err := dbServer.CreateReply(reply)
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

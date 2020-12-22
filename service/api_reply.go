package service

import (
	"net/http"
)

func CreateReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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
	replies, err := dbServer.GetReplies(int64(article_id))
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

	msg, err := json.Marshal(replies)
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
	_, err := dbServer.LikeReply(int64(reply_id), int64(1))
	if err != nil {
		log.Fatal("Fail to update reply's like number by ID", err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	w.WriteHeader(http.StatusOK)
}

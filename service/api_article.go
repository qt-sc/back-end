package service

import (
	"encoding/json"
	"github.com/qt-sc/server/lib"
	"log"
	"net/http"
	"strconv"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "users")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
	}
	articles, err := dbServer.GetArticleByArticle(int64(id))
	if err != nil {
		log.Fatal("Fail to get article by ID", err)
	}

	msg, err := json.Marshal(articles)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
	}
	w.Write(msg)
}

func GetArticlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "users")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
	}
	articles, err := dbServer.GetArticleByUser(int64(id))
	if err != nil {
		log.Fatal("Fail to get article by ID", err)
	}

	msg, err := json.Marshal(articles)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
	}
	w.Write(msg)
}


func GetArticlesPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetCreateArticlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func LikeArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

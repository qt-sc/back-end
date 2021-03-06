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

func GetAllArticles(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	articleList, err := dbServer.GetAllArticle()
	if err != nil {
		log.Fatal("Fail to get all article", err)
		w.WriteHeader(http.StatusNotFound)
	}

	msg, err := json.Marshal(articleList)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)

	w.WriteHeader(http.StatusOK)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	body, _ := ioutil.ReadAll(r.Body)

	var article model.Article
	json.Unmarshal([]byte(body), &article)

	_, err := dbServer.CreateArticle(article)
	if err != nil {
		log.Fatal("Fail to create article", err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	w.WriteHeader(http.StatusOK)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "articles")
	article_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}

	_, err = dbServer.DeleteArticle(int64(article_id))
	if err != nil {
		log.Fatal("Fail to delete article by ID", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "articles")
	article_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}
	article, err := dbServer.GetArticleByArticle(int64(article_id))
	if err != nil {
		log.Fatal("Fail to get article by ID", err)
		w.WriteHeader(http.StatusNotFound)
	}

	msg, err := json.Marshal(article)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)

	w.WriteHeader(http.StatusOK)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "users")
	user_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}
	articles, err := dbServer.GetArticleByUser(int64(user_id))
	if err != nil {
		log.Fatal("Fail to get article by ID", err)
		w.WriteHeader(http.StatusNotFound)
	}

	msg, err := json.Marshal(articles)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)
}

func LikeArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "articles")
	article_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}
	_, err = dbServer.UpdateArticleLikeNum(int64(article_id), int64(1))
	if err != nil {
		log.Fatal("Fail to update article's like number by ID", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	body, _ := ioutil.ReadAll(r.Body)

	var article model.Article
	json.Unmarshal([]byte(body), &article)

	_, err := dbServer.UpdateArticleContent(article.Id, article.Content)
	if err != nil {
		log.Fatal("Fail to update article content", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

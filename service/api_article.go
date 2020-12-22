package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/qt-sc/server/lib"
	"github.com/qt-sc/server/model"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	r.ParseForm()
	idstr := r.PostFormValue("id")
	article_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}

	article_readNum, err :=  strconv.Atoi(r.PostFormValue("readNum"))
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}

	article_likeNum, err :=  strconv.Atoi(r.PostFormValue("likeNum"))
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}

	article_user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
	
	article := model.Article{
		Id:       	int64(article_id),
		Title:		r.PostFormValue("title"),
		ReadNum:	int64(article_readNum),
		LikeNum: 	int64(article_likeNum),
		Content:	r.PostFormValue("content"),
		UserID:		int64(article_user_id),
		Replies:	nil,
		Tags:		nil,
		Url:      	r.PostFormValue("url"),
	}

	_, err = dbServer.CreateArticle(article)
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

// func GetArticlePage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// }

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

// func GetArticlesPage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// }

// func GetCreateArticlePage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// }

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
	
	r.ParseForm()
	idstr := r.PostFormValue("id")
	article_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}

	_, err = dbServer.UpdateArticleContent(int64(article_id), r.PostFormValue("content"))
	if err != nil {
		log.Fatal("Fail to update article content", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}

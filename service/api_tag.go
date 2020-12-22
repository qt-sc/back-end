package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/qt-sc/server/lib"
	"github.com/qt-sc/server/model"
)

// func GetArticlesPageByTag(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// }

func CreateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	body, _ := ioutil.ReadAll(r.Body)

	var tag model.Tag
	json.Unmarshal([]byte(body), &tag)

	_, err := dbServer.CreateTag(tag)
	if err != nil {
		log.Fatal("Fail to create tag", err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	w.WriteHeader(http.StatusOK)
}

func GetTags(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	tags, err := dbServer.GetAllTag()
	if err != nil {
		log.Fatal("Fail to get all tags", err)
		w.WriteHeader(http.StatusNotFound)
	}

	msg, err := json.Marshal(tags)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)

	w.WriteHeader(http.StatusOK)
}

func SearchArticlesByTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	url := r.RequestURI
	name := lib.GetFollowParameter(url, "tags")

	articles, err := dbServer.GetArticleByTag(name)
	if err != nil {
		log.Fatal("Fail to get article by tag name", err)
		w.WriteHeader(http.StatusNotFound)
	}

	msg, err := json.Marshal(articles)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)

	w.WriteHeader(http.StatusOK)
}

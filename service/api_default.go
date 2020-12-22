package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/qt-sc/server/model"
)

func GetApis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	apis := model.Apis{
		UsersUrl:    "https://localhost:8080/api/users",
		UserUrl:     "https://localhost:8080/api/users/{userid}",
		ArticlesUrl: "https://localhost:8080/api/users/{userid}/articles",
		ArticleUrl:  "https://localhost:8080/api/users/{userid}/articles/{articleid}",
		RepliesUrl:  "https://localhost:8080/api/users/{userid}/articles/{articleid}/replies",
		ReplyUrl:    "https://localhost:8080/api/users/{userid}/articles/{articleid}/replies/{replyid}",
		TagsUrl:     "https://localhost:8080/api/tags",
		TagUrl:      "https://localhost:8080/api/tags/{tagname}",
	}

	msg, err := json.Marshal(apis)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)

	w.WriteHeader(http.StatusOK)
}

func ApiGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	apis := model.Apis{
		UsersUrl: "https://localhost:8080/api/users",
		UserUrl: "https://localhost:8080/api/users/{userid}",
		ArticlesUrl: "https://localhost:8080/api/users/{userid}/articles",
		ArticleUrl: "https://localhost:8080/api/users/{userid}/articles/{articleid}",
		RepliesUrl: "https://localhost:8080/api/users/{userid}/articles/{articleid}/replies",
		ReplyUrl: "https://localhost:8080/api/users/{userid}/articles/{articleid}/replies/{replyid}",
		TagsUrl: "https://localhost:8080/api/tags",
		TagUrl: "https://localhost:8080/api/tags/{tagname}",
	}

	msg, err := json.Marshal(apis)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)
	
	w.WriteHeader(http.StatusOK)
}
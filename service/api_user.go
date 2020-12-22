package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/qt-sc/server/conf"
	"github.com/qt-sc/server/lib"
	"github.com/qt-sc/server/model"
)

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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	url := r.RequestURI
	idstr := lib.GetFollowParameter(url, "users")
	user_id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Fatal("string to int fail", err)
		w.WriteHeader(http.StatusNotFound)
	}

	_, err = dbServer.DeleteUser(int64(user_id))
	if err != nil {
		log.Fatal("Fail to delete user by ID", err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	w.WriteHeader(http.StatusOK)
}

// func GetSignupPage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// }

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	username, err := r.Cookie("username")
	if err != nil {
		log.Println("获取cookie失败")
	}

	user, err := dbServer.GetOneUser(username.Value)
	if err != nil {
		log.Fatal("Fail to get user by ID", err)
		w.WriteHeader(http.StatusNotFound)
	}

	msg, err := json.Marshal(user)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(msg)
	
	w.WriteHeader(http.StatusOK)
}

// func GetUserPage(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// }

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	users, err := dbServer.GetAllUser()
	if err != nil {
		log.Fatal("Fail to get all users", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	msg, err := json.Marshal(users)
	if err != nil {
		log.Fatal("JSON Marshal fail.", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write(msg)

	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	body, _ := ioutil.ReadAll(r.Body)

	var user model.User
	json.Unmarshal([]byte(body), &user)

	_, err := dbServer.UpdateUser(user)
	if err != nil {
		log.Fatal("Fail to update user", err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	w.WriteHeader(http.StatusOK)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	// TODO：登录相关的非鉴权逻辑
	r.ParseForm()
	id := r.PostFormValue("id")
	username := r.PostFormValue("name")
	user, err := dbServer.GetOneUser(username)
	if err != nil {
		log.Println("获取用户失败")
		return
	}
	password := r.PostFormValue("password")
	if password != user.Password {
		log.Println("密码错误，登录失败")
		return
	}

	// TODO： 登录相关的鉴权逻辑
	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("类型转换错误")
	}

	userTemp := lib.UserInfo{
		Username: username,
		ID:       uint64(userId),
	}

	token, err := lib.CreateToken(&userTemp)

	if err != nil {
		log.Println("token生成错误")
	}

	conf.Redis.Set(token, userId, time.Hour*24)

	cookie1 := http.Cookie{
		Name: "username",
		Value:username,
		Path:"/",
		Expires:time.Now().AddDate(0, 0, 1),
	}

	cookie2 := http.Cookie{
		Name:       "id",
		Value:      id,
		Path:"/",
		Expires:    time.Now().AddDate(0, 0, 1),
	}

	cookie3 := http.Cookie{
		Name:"token",
		Value:token,
		Path:"/",
		Expires:    time.Now().AddDate(0, 0, 1),
	}

	w.Header().Set("Set-Cookie", cookie1.String())
	w.Header().Add("Set-Cookie", cookie2.String())
	w.Header().Add("Set-Cookie", cookie3.String())

}

func UserLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	token, err := r.Cookie("token")
	if err != nil {
		log.Println("获取cookie失败")
		return
	}

	conf.Redis.Del(token.Value) 
	w.Header().Del("Cookie")
}

func UserSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	
	body, _ := ioutil.ReadAll(r.Body)

	var user model.User
	json.Unmarshal([]byte(body), &user)

	_, err := dbServer.CreateUser(user)
	if err != nil {
		log.Fatal("Fail to create user", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
}
package service

import (
	"fmt"
	"github.com/qt-sc/server/database"
	"github.com/qt-sc/server/script"
)

var dbServer database.DBServiceInterface

//var conn redis.Conn
//var hubServer *noti.Hub

func init() {
	dbServer = &database.DBService{}
	//var err error
	//conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	log.Println("连接到rpc服务器失败")
	//	return
	//}

	//hubServer = noti.NewHubInstance()
	//go hubServer.Run()

	getZhihuDaily()
	fmt.Println(dbServer.GetArticleByArticle(1))
}

func getZhihuDaily() {
	ref := script.GetLatestEssay()

	fmt.Println(ref[0])

	for _, x := range ref {
		ok, err := dbServer.CreateArticle(x)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !ok {
			continue
		}
	}
}

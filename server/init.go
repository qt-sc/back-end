package server

import (
	"github.com/qt-sc/server/database"
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

	getLatestDaily()
}

func getLatestDaily()  {

}
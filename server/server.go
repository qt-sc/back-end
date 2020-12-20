package server

import (
	"github.com/qt-sc/server/lib"
	"github.com/qt-sc/server/model"
	"log"
	"net/http"
	"time"
	"io/ioutil"

	"encoding/json"
)

//CreateEssayRequest 用户创建博客请求
type CreateEssayRequest struct {
	//IsPublic    string `json:"ispublic"`
	Content     string `json:"content"`
	//PictureName string `json:"picture_name"`
}

//PublicEssaysResponse 返回用户所有可见博客
type PublicEssaysResponse struct {
	ID          string `json:"id"`
	CreateTime  string `json:"create_time"`
	Content     string `json:"content"`
	//CreatorName string `json:"creator_name"`
	//PictureName string `json:"picture_name"`
	//GoodCount   int    `json:"good_count"`
}

//GetAllEssayPublic 获取当前所有public博客
func GetAllEssayPublic(w http.ResponseWriter, req *http.Request) (bool, interface{}) {
	allEssayData, err := dbServer.GetAllPublicEssay()
	if err != nil {
		log.Println(err)
		return false, "获取所有公开博客失败"
	}
	var publicEssayResponseList []PublicEssaysResponse

	for _, v := range allEssayData {
		//user, err := dbServer.GetUserFromEmail(v.CreatorEmail)
		//if err != nil {
		//	return false, "获取所有公开博客失败"
		//}
		publicEssayResponseList = append(publicEssayResponseList, PublicEssaysResponse{
			ID:          v.ID,
			CreateTime:  v.CreateTime,
			Content:     v.Content,
			//CreatorName: user.Username,
			//GoodCount:   v.GoodCount,
			//PictureName: v.PictureName,
		})
	}
	return true, publicEssayResponseList
}

//CreateEssayHandler 提供创建博客服务
func CreateEssayHandler(w http.ResponseWriter, req *http.Request) (bool, interface{}) {
	//vars := mux.Vars(req)
	//useremail := vars["email"]
	// Check token
	//ok, Tuseremail := lib.GetUserEmailFromToken(req.Header.Get("token"), lib.SignKey)
	//if !ok {
	//	return false, "身份验证失败"
	//}
	//if Tuseremail != useremail {
	//	return false, "身份验证失败"
	//}

	//if _, err := dbServer.GetUserFromEmail(useremail); err == gorm.ErrRecordNotFound {
	//	return false, "用户不存在"
	//}

	crateEssayRequest := CreateEssayRequest{}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return false, "无法读取用户的请求"
	}
	if err := json.Unmarshal(body, &crateEssayRequest); err != nil {
		log.Println(err)
		return false, "无效的json信息"
	}

	if crateEssayRequest.Content == "" {
		log.Print("博客内容不能为空")
		return false, "博客内容不能为空"
	}
	essay := model.Essay{
		//CreatorEmail: useremail,
		CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
		ID:           lib.GetUniqueID(),
		//IsPublic:     crateEssayRequest.IsPublic,
		Content:      crateEssayRequest.Content,
		//PictureName:  crateEssayRequest.PictureName,
		//GoodCount:    0,
	}

	if ok, err := dbServer.UserCreateEssay(essay); ok != true || err != nil {
		log.Println(err)
		return false, "创建博客失败"
	}
	return true, ""
}
package script

import (
	"encoding/json"
	"fmt"
	"github.com/qt-sc/server/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var ZHIHU_URL string = "http://news-at.zhihu.com/api/4"

type DailyStory struct {
	Title     string   `json:"title"`
	Ga_prefix string   `json:"ga_prefix"`
	Images    []string `json:"images"`
	Types     int      `json:"type"`
	Id        int      `json:"id"`
}

type DailyList struct {
	Date        string       `json:"date"`
	Stories     []DailyStory `json:"stories"`
	Top_stories []DailyStory `json:"top_stories"`
}

type DailyEssay struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

type DailyExtra struct {
	Popularity int `json:"popularity"`
}

type DailyComment struct {
	Comments []DailyReply `json:"comments"`
}

type DailyReply struct {
	Author  string `json:"author"`
	Id      int    `json:"id"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
	Time    int    `json:"time"`
}

func getLatestEssay() []model.Article {
	dailyIDList, err := getDailyList()
	if err != nil {
		fmt.Println("Fail to get Daily Essay List.")
		return nil
	}

	var articleArr []model.Article

	for i, id := range dailyIDList {
		article := model.Article{}
		article.Title, article.Content, err = getEssayById(id)
		if err != nil {
			fmt.Println("Fail to get content.")
			continue
		}
		reply, err := getReplyById(id)
		if err != nil {
			fmt.Println("Fail to get reply.")
			continue
		}

		for _, x := range reply {
			mid := model.Reply{int64(x.Id), int64(x.Likes), time.Now(), x.Content, "/users/" + x.Author}
			article.Replies = append(article.Replies, mid)
		}

		article.LikeNum, err = int64(getExtraById(id))
		if err != nil {
			fmt.Println("Fail to get likenum.")
			continue
		}
		article.ReadNum = 0
		article.Id = int64(i)

		articleArr = append(articleArr, article)
	}

	return articleArr
}

func getDailyList() ([]int, error) {
	rsp, err := http.Get(ZHIHU_URL + "/news/latest")
	if err != nil {
		fmt.Println("get Zhihu Daily has some error. Error info: ", err)
		return []int{}, err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("get Zhihu Daily has some error. Error info: ", err)
		return []int{}, err
	}

	rspList := &DailyList{}
	err = json.Unmarshal(body, rspList)
	if err != nil {
		fmt.Println("get Zhihu Daily has some error. Error info: ", err)
		return []int{}, err
	}

	var dailyIDList []int
	for _, x := range rspList.Stories {
		dailyIDList = append(dailyIDList, x.Id)
	}
	for _, x := range rspList.Top_stories {
		dailyIDList = append(dailyIDList, x.Id)
	}

	return dailyIDList, nil
}

func getEssayById(id int) (string, string, error) {
	rsp, err := http.Get(ZHIHU_URL + "/news/" + strconv.Itoa(id))
	if err != nil {
		fmt.Println("get Zhihu Essay has some error. Error info: ", err)
		return "", "", err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("get Zhihu Essay has some error. Error info: ", err)
		return "", "", err
	}

	rspEssay := &DailyEssay{}
	err = json.Unmarshal(body, rspEssay)
	if err != nil {
		fmt.Println("get Zhihu Essay has some error. Error info: ", err)
		return "", "", err
	}

	//fmt.Println(rspEssay.Body)
	return rspEssay.Title, rspEssay.Body, nil
}

func getReplyById(id int) ([]DailyReply, error) {
	rsp, err := http.Get(ZHIHU_URL + "/story/" + strconv.Itoa(id) + "/short-comments")
	if err != nil {
		fmt.Println("get Zhihu Essay Reply has some error. Error info: ", err)
		return nil, err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("get Zhihu Essay Reply has some error. Error info: ", err)
		return nil, err
	}

	rspComment := &DailyComment{}
	err = json.Unmarshal(body, rspComment)
	if err != nil {
		fmt.Println("get Zhihu Essay Reply has some error. Error info: ", err)
		return nil, err
	}

	//fmt.Println(rspComment.Body)
	return rspComment.Comments, nil
}

func getExtraById(id int) (int, error) {
	rsp, err := http.Get(ZHIHU_URL + "/story-extra/" + strconv.Itoa(id))
	if err != nil {
		fmt.Println("get Zhihu Essay Extra Info has some error. Error info: ", err)
		return -1, err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("get Zhihu Essay Extra Info has some error. Error info: ", err)
		return -1, err
	}

	rspExtra := &DailyExtra{}
	err = json.Unmarshal(body, rspExtra)
	if err != nil {
		fmt.Println("get Zhihu Essay Extra Info has some error. Error info: ", err)
		return -1, err
	}

	//fmt.Println(rspEssay.Body)
	return rspExtra.Popularity, nil
}

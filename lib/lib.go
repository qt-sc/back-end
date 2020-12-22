package lib

import (
	"github.com/rs/xid"
	"log"
	"regexp"
	"strings"
)

//GetUniqueID 返回唯一哈希值
func GetUniqueID() string {
	return xid.New().String()
}

func GetFollowParameter(url string, key string) string {
	pa := strings.Split(url, "/")
	var value string = ""
	for i, x := range pa {
		if x == key && i < len(pa)-1 {
			value = pa[i+1]
		}
	}
	return value
}

func RemoveNonBmpUnicode(str string) string {
	re, err := regexp.Compile("[^\\x00-\\xFF]")
	if err != nil {
		log.Fatal(err)
	}
	rep := re.ReplaceAllString(str, "")
	return rep
}
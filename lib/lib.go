package lib

import (
	"github.com/rs/xid"
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

//func RemoveNonBmpUnicode(str string) string {
//	re, err := regexp.Compile("[^\\x00-\\xFF]")
//	if err != nil {
//		log.Fatal(err)
//	}
//	rep := re.ReplaceAllString(str, "")
//	return rep
//}

func RemoveNonBmpUnicode(str string) string {
	ref := []byte(str)

	for i:=0; i< len(ref); i++ {
		if ((ref[i] & 0xF8) == 0xF0) {
			for j:=0; j<4; j++ {
				ref[i+j] = 0x3F
			}
		}
		i+=3
	}

	return string(ref[:])
}
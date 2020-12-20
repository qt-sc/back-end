package lib

import (
	"github.com/rs/xid"
)

//GetUniqueID 返回唯一哈希值
func GetUniqueID() string {
	return xid.New().String()
}

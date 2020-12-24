package middleware

import (
	"github.com/qt-sc/server/conf"
	"github.com/qt-sc/server/lib"
	"log"
	"net/http"
)

func Auth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// TODO： 如果是login和signup，不需要鉴权
		// 若有其他无需鉴权的api，请添加
		// 为了做测试把/api/前缀api剔除
		// if r.URL.Path == "/users/login" || r.URL.Path == "/users/signup" || strings.HasPrefix(r.URL.Path, "/api/"){
		if r.URL.Path == "/users/login" || r.URL.Path == "/users/signup"){
			next.ServeHTTP(w, r)
			return
		}

		token, err := r.Cookie("token")
		if err != nil {
			log.Println("token获取失败")
			return
		}
		err = lib.ParseToken(token.Value)
		if err != nil {
			log.Println("鉴权失败，没有访问权限")
			return
		}

		if conf.Redis.Get(token.Value).Val() == "" {
			log.Println("鉴权失败, 没有访问权限")
			return
		}

		log.Println("鉴权有效")
		next.ServeHTTP(w, r)
	})
}


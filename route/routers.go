package route

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/qt-sc/server/service"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreateArticle",
		strings.ToUpper("Post"),
		"/users/{userid}/articles/new",
		service.CreateArticle,
	},

	Route{
		"DeleteArticle",
		strings.ToUpper("Delete"),
		"/users/{userid}/articles/{articleid}",
		service.DeleteArticle,
	},

	Route{
		"GetArticle",
		strings.ToUpper("Get"),
		"/api/users/{userid}/articles/{articleid}",
		service.GetArticle,
	},

	Route{
		"GetArticles",
		strings.ToUpper("Get"),
		"/api/users/{userid}/articles",
		service.GetArticles,
	},

	Route{
		"LikeArticle",
		strings.ToUpper("Put"),
		"/users/{userid}/articles/{articleid}/like",
		service.LikeArticle,
	},

	Route{
		"UpdateArticle",
		strings.ToUpper("Put"),
		"/users/{userid}/articles/{articleid}",
		service.UpdateArticle,
	},

	Route{
		"GetApis",
		strings.ToUpper("Get"),
		"/api",
		service.GetApis,
	},

	Route{
		"CreateReply",
		strings.ToUpper("Post"),
		"/users/{userid}/articles/{articleid}/replies/new",
		service.CreateReply,
	},

	Route{
		"GetReplies",
		strings.ToUpper("Get"),
		"/api/users/{userid}/articles/{articleid}/replies",
		service.GetReplies,
	},

	Route{
		"GetReply",
		strings.ToUpper("Get"),
		"/api/users/{userid}/articles/{articleid}/replies/{replyid}",
		service.GetReply,
	},

	Route{
		"LikeReply",
		strings.ToUpper("Put"),
		"/users/{userid}/articles/{articleid}/replies/{replyid}/like",
		service.LikeReply,
	},

	Route{
		"CreateTag",
		strings.ToUpper("Post"),
		"/tags/new",
		service.CreateTag,
	},

	Route{
		"GetTags",
		strings.ToUpper("Get"),
		"/api/tags",
		service.GetTags,
	},

	Route{
		"SearchArticlesByTag",
		strings.ToUpper("Get"),
		"/api/tags/{tagname}/articles",
		service.SearchArticlesByTag,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/users/{userid}",
		service.DeleteUser,
	},

	Route{
		"GetUserInfo",
		strings.ToUpper("Get"),
		"/api/users/{userid}",
		service.GetUserInfo,
	},

	Route{
		"GetUsers",
		strings.ToUpper("Get"),
		"/api/users",
		service.GetUsers,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/users/{userid}",
		service.UpdateUser,
	},

	Route{
		"UserLogin",
		strings.ToUpper("Post"),
		"/users/login",
		service.UserLogin,
	},

	Route{
		"UserLogout",
		strings.ToUpper("Get"),
		"/users/logout",
		service.UserLogout,
	},

	Route{
		"UserSignup",
		strings.ToUpper("Post"),
		"/users/signup",
		service.UserSignup,
	},
}

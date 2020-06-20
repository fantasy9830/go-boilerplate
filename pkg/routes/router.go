package routes

import (
	"go-boilerplate/pkg/auth"
	"go-boilerplate/pkg/user"

	"github.com/gin-gonic/gin"
)

var _ IRouter = (*Router)(nil)

// IRouter IRouter
type IRouter interface {
	RegisterAPI(api *gin.RouterGroup) error
	RegisterWeb(web *gin.RouterGroup) error
}

// Router 路由註冊
type Router struct {
	Auth *auth.Controller
	User *user.Controller
}

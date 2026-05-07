package routers

import (
	"github.com/haidaqn/go-ecommerce-backend-api/internal/router/admin"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/router/auth"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/router/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin admin.AdminRouterGroup
	Auth  auth.AuthRouterGroup
}

var RouterGroupApp = new(RouterGroup)

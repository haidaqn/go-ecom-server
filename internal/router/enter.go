package routers

import (
	"github.com/haidaqn/go-ecommerce-backend-api/internal/router/auth"
)

type RouterGroup struct {
	Auth auth.AuthRouterGroup
}

var RouterGroupApp = new(RouterGroup)

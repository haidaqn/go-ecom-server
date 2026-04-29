package main

import (
	routers "github.com/haidaqn/go-ecommerce-backend-api/internal/router"
)

func main() {
	r := routers.NewRoute()

	r.Run(":3000")
}

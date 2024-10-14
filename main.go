package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/khanhuyy/understream/component"
	"github.com/khanhuyy/understream/router"
	"log"
)

func ServeHTTPRouter() {
	l := &log.Logger{}
	component.BuildConfig(l)
	port := fmt.Sprintf(":%d", 8080)

	ctx := &gin.Context{}

	r := gin.Default()
	router.BuildRouter(r, ctx)
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}

func main() {
	ServeHTTPRouter()
}

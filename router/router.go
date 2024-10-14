package router

import (
	"github.com/gin-gonic/gin"
	"github.com/khanhuyy/understream/internal/handler"
)

func BuildRouter(router *gin.Engine, ctx *gin.Context) *gin.Engine {
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/", handler.GetDashboard(ctx))
	router.GET("/videos", handler.GetListVideo(ctx))
	//router.GET("/videos/:id/stream")
	router.GET("/videos/:filename", handler.Stream(ctx))
	router.GET("/users")
	router.POST("/login", handler.Login())
	router.POST("/register", handler.Register())
	return router
}

package handler

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func GetDashboard(*gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello world": "this is base dashboard",
		})
	}
}

func GetListVideo(*gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"url": "youtube.com",
		})
	}
}

func Stream(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Param("filename")
		file, err := os.Open("videos/" + filename)
		if err != nil {
			c.String(http.StatusNotFound, "Video not found.")
			return
		}
		defer file.Close()

		c.Header("Content-Type", "video/mp4")
		buffer := make([]byte, 64*1024) // 64KB buffer size
		io.CopyBuffer(c.Writer, file, buffer)
	}
}

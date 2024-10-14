package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/khanhuyy/understream/component/database"
	"github.com/khanhuyy/understream/internal/biz"
	"github.com/khanhuyy/understream/internal/repository"
	"net/http"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.GetPostForm("email")
		password, _ := c.GetPostForm("password")
		db := database.InitGORM()
		service := biz.NewUserService(repository.NewSQLStore(db))
		user, err := service.Login(email, password)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":       user.Id,
			"email":    user.Email,
			"username": user.Username,
		})
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.GetPostForm("email")
		password, _ := c.GetPostForm("password")
		confirmPassword, _ := c.GetPostForm("confirm_password")
		db := database.InitGORM()
		service := biz.NewUserService(repository.NewSQLStore(db))
		user, err := service.Register(email, password, confirmPassword)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":       user.Id,
			"email":    user.Email,
			"username": user.Username,
		})
	}
}

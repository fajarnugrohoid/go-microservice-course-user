package middleware

import (
	"context"
	"course/internal/domain"
	"course/internal/user/usecase"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithAuth(userUsecase *usecase.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		auths := strings.Split(authHeader, " ")
		if len(auths) != 2 {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		user := domain.User{}
		data, err := user.DecryptJWT(auths[1])
		if err != nil {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		userId := int(data["user_id"].(float64))

		log.Println("WithAuth userId:", userId)

		dbUser, err := userUsecase.GetUserByID(c.Request.Context(), userId)
		log.Println("err:", err)
		if err != nil || dbUser.ID == 0 {
			c.JSON(401, map[string]string{
				"message": "unathorized",
			})
			c.Abort()
			return
		}
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", userId)
		c.Request = c.Request.WithContext(ctxUserID)
		c.Next()
	}
}

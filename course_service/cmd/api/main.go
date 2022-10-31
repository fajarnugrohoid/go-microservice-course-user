package main

import (
	"course/internal/database"
	"course/internal/exercise"
	"course/internal/middleware"
	"course/internal/user/repository"
	"course/internal/user/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("userMscv")
	route := gin.Default()

	route.GET("/hello", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"message": "hello world",
		})
	})

	dbConn := database.NewDatabaseConn()
	eu := exercise.NewExerciseUsecase(dbConn)
	//userRepo := repository.NewUserDBRepo(dbConn)
	userMscv := repository.NewMcsUserRepo()

	fmt.Println("userMscv:", userMscv)
	log.Println("userMscv.Log:", userMscv)

	//uu := usecase.NewUserUsecase(userRepo)
	uu := usecase.NewUserUsecase(userMscv)

	// usecase endpoint
	route.GET("/exercises/:id", middleware.WithAuth(uu), eu.GetExerciseByID)
	route.GET("/exercises/:id/scores", middleware.WithAuth(uu), eu.GetScore)

	// user endpoint
	//route.POST("/register", uu.Register)
	//route.POST("/login", uu.Login)

	route.Run(":1234")
}

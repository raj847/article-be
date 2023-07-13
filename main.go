package main

import (
	"article/handler"
	"article/repository"
	"article/service"
	"article/utils"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := os.Setenv("DATABASE_URL", "root:@tcp(127.0.0.1:3306)/article?parseTime=true")
	if err != nil {
		log.Fatalf("cannot set env: %v", err)
	}
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	// Routes
	e.POST("/article", postHandler.Create)
	e.PATCH("/article/:id/trash", postHandler.UpdateTrash)
	e.GET("/article/:limit/:offset", postHandler.GetAllPagination)
	e.GET("/article/:id", postHandler.GetByID)
	e.PUT("/article/:id", postHandler.Update)
	e.DELETE("/article/:id", postHandler.Delete)
	e.GET("/article", postHandler.GetAll)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

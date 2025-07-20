package router

import (
	"quiz3/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)

	router.POST("/api/categories", controllers.PostCategory)
	router.GET("/api/categories", controllers.GetCategories)
	router.GET("/api/categories/:id", controllers.GetCategoryById)
	router.GET("/api/categories/:id/books", controllers.GetBooksByCategoryId)
	router.PUT("/api/categories/:id", controllers.UpdateCategory)
	router.DELETE("/api/categories/:id", controllers.DeleteCategory)

	router.POST("/api/books", controllers.PostBook)
	router.GET("/api/books", controllers.GetBooks)
	router.GET("/api/books/:id", controllers.GetBookById)
	router.PUT("/api/books/:id", controllers.UpdateBook)
	router.DELETE("/api/books/:id", controllers.DeleteBook)

	return router
}

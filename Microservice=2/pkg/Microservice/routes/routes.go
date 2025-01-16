package routes

import (
	"api-gateway/pkg/Microservice/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	userRoutes := router.Group("/users") // Group routes under `/users`
	{
		userRoutes.POST("/method", userHandler.Methods)
		userRoutes.POST("/create", userHandler.CreateUser)       // POST: Create a new user
		userRoutes.GET("/:id", userHandler.GetUserByID)          // GET: Get user by ID
		userRoutes.GET("/get-all", userHandler.ListAllUsers)     // GET: List all users
		userRoutes.PATCH("/update", userHandler.UpdateUser)      // PATCH: Update user details
		userRoutes.DELETE("/delete/:id", userHandler.DeleteUser) // DELETE: Delete user by ID
	}
}

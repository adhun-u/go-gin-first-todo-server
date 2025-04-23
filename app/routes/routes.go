package routes

import (
	"func/app/controllers"
	"github.com/gin-gonic/gin"
)

// Get todos route
func Get_todo_route(route *gin.Engine) {
	route.GET("/get/todos", controllers.Get_todos_controller)
}

// Post todo route
func Post_todo_route(route *gin.Engine) {
	route.POST("/post/todo", controllers.Post_todo_controller)
}

// Delete todo route
func Delete_todo_route(route *gin.Engine) {
	route.DELETE("delete/todo", controllers.Delete_todo_controller)
}

// Update todo route
func Update_todo_route(route *gin.Engine) {
	route.PUT("/update/todo", controllers.Update_todo_controller)
}

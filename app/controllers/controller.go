package controllers

import (
	"func/app/models"
	"net/http"
	"slices"
	"github.com/gin-gonic/gin"
)

var todo = []models.LoginInfo{}

// Get controller
func Get_todos_controller(ctx *gin.Context) {
	if len(todo) == 0 {
		ctx.JSON(404, gin.H{
			"status_code": 404,
			"message":     "there are no tasks",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":     "success",
		"status_code": 200,
		"tasks":       todo,
	})
}

// Post controller
func Post_todo_controller(ctx *gin.Context) {
	var input models.LoginInfo
	error := ctx.ShouldBindJSON(&input)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messaage": "Invalid json"})
		return
	}
	newTodo := models.LoginInfo{
		Id:          input.Id,
		Task:        input.Task,
		Description: input.Description,
	}
	todo = append(todo, newTodo)
	ctx.JSON(200, gin.H{
		"message": "Added successfully",
	})
}

// Delete controller
func Delete_todo_controller(ctx *gin.Context) {
	var Id models.Idonly
	error := ctx.ShouldBindJSON(&Id)

	if error != nil {
		ctx.JSON(404, gin.H{
			"message": error,
		})
		return
	}
	index := slices.IndexFunc(todo, func(todo models.LoginInfo) bool {
		return todo.Id == Id.Id
	})
	todo = slices.Delete(todo, index, index+1)
	ctx.JSON(200, gin.H{
		"message": "Deleted successfully",
	})
}

// Update controller
func Update_todo_controller(ctx *gin.Context) {
	var updateDetails models.LoginInfo
	error := ctx.ShouldBindJSON(&updateDetails)
	if error != nil {
		ctx.JSON(404, gin.H{"message": "Invalid json"})
		return
	}
	index := slices.IndexFunc(todo, func(item models.LoginInfo) bool {
		return item.Id == updateDetails.Id
	})
	details := models.LoginInfo{
		Id:          updateDetails.Id,
		Task:        updateDetails.Task,
		Description: updateDetails.Description,
	}
	todo = slices.Replace(todo, index, index+1, details)
	ctx.JSON(200, gin.H{"message": "updated successfully"})
}

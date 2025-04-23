package main

import (
	"fmt"
	"func/app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println()
	route := gin.Default()
	routes.Get_todo_route(route)
	routes.Post_todo_route(route)
	routes.Delete_todo_route(route)
	routes.Update_todo_route(route)
	route.Run()
}

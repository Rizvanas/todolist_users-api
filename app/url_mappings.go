package app

import (
	"github.com/Rizvanas/todolist_users-api/controllers/ping"
	"github.com/Rizvanas/todolist_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUserByID)
	router.POST("/users", users.CreateUser)
}

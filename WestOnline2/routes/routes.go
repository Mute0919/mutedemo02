package routes

import (
	"WestOnline2/api"
	"WestOnline2/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("tasks", api.ListTask)
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.DELETE("task/:id", api.DeleteTask)
			authed.PUT("task/:id", api.UpdateTask)
			authed.POST("search", api.SearchTasks)
		}
	}
	return r
}

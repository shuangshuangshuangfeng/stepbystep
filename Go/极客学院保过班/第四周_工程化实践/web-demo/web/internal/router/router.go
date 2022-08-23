package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-demo/web/internal/router/middleware"
	"web-demo/web/internal/service"
)

//InitRouter 初始化路由
func InitRouter(g *gin.Engine) {
	middlewares := []gin.HandlerFunc{}
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middlewares...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})


	// The health check handlers
	router := g.Group("/user")
	{
		router.POST("/addUser", service.AddUser)                    //添加用户
		router.POST("/selectUser", service.SelectUser)          //查询用户
		router.GET("/sayHiUser", service.SayHiUser)          //查询用户  http://127.0.0.1:8080/user/sayHiUser
	}

}


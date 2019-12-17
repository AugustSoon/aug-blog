package router

import (
	"github.com/JumpSama/aug-blog/handler/user"
	"github.com/JumpSama/aug-blog/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 中间件
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 登录
	g.POST("/api/login", user.Login)

	// 用户
	u := g.Group("/api/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.GET("/:account", user.Get)
		u.GET("", user.List)
		u.PUT("/:id", user.Update)
	}

	return g
}

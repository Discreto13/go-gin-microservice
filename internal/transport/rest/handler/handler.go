package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	user *UserHandler
}

func NewHandler(userHandler *UserHandler) *Handler {
	return &Handler{
		user: userHandler,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.registerAPI(router)

	return router
}

func (h *Handler) registerAPI(router *gin.Engine) {
	api := router.Group("/api")
	{
		users := api.Group("/user")
		{
			users.POST("", h.user.Create)
			users.GET(":id", h.user.GetById)
			users.GET("", h.user.GetAll)
		}
	}
}

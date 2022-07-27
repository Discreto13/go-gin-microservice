package handler

import (
	"context"
	"net/http"

	"github.com/discreto13/go-gin-microservice/internal/core"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(ctx context.Context, user *core.CreateUser) (*core.User, error)
	GetById(ctx context.Context, id string) (*core.User, error)
	GetAll(ctx context.Context) ([]*core.User, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	var dto core.CreateUser
	if err := c.BindJSON(&dto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createdUser, err := h.service.Create(c, &dto)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetById(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	deleted, err := h.service.Delete(c, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !deleted {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "nothing to delete"})
		return
	}

	c.IndentedJSON(http.StatusOK,  gin.H{"message": "success"})
}

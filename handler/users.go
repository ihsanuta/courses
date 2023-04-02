package handler

import (
	"courses/app/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) Login(c *gin.Context) {
	var form model.Login

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.usecase.User.Login(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    result,
	})
}

func (h *handler) Register(c *gin.Context) {
	var form model.User

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.User.Register(&form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    form,
	})
}

func (h *handler) DeleteUser(c *gin.Context) {
	paramId := c.Param("id")
	if paramId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter id not found"})
		return
	}

	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter id must integer"})
		return
	}

	err = h.usecase.User.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
	})
}

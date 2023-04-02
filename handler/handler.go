package handler

import (
	"courses/app/model"
	"courses/app/usecase"
	"courses/config"
	"courses/helper/token"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type Handler interface{}

var once = &sync.Once{}

type handler struct {
	usecase *usecase.Usecase
}

func Init(usecase *usecase.Usecase) *handler {
	var h *handler
	once.Do(func() {
		h = &handler{
			usecase: usecase,
		}
		h.Serve()
	})
	return h
}

func (h *handler) Serve() {
	router := gin.Default()
	group := router.Group("/api/v1")
	group.DELETE("/admin/user/:id", h.authenticateAdminToken, h.DeleteUser)
	group.POST("/admin/login", h.Login)
	group.POST("/admin/course", h.authenticateAdminToken, h.InsertCourse)
	group.PUT("/admin/course/:id", h.authenticateAdminToken, h.UpdateCourse)
	group.GET("/admin/course", h.authenticateAdminToken, h.GetListCourse)
	group.GET("/admin/course/:id", h.authenticateAdminToken, h.GetCourse)
	group.GET("/admin/statistic", h.authenticateAdminToken, h.GetStatistic)

	group.POST("/user/register", h.Register)
	group.GET("/user/course", h.GetListCourse)
	group.GET("/user/course/:id", h.GetCourse)
	group.GET("/user/category/course", h.GetCategory)

	serverString := fmt.Sprintf("%s:%s", config.AppConfig["host"], config.AppConfig["port"])
	router.Run(serverString)
}

func (h *handler) authenticateAdminToken(c *gin.Context) {
	tokenString := GetTokenFromGinContext(c)

	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		c.Abort()
		return
	}
	data, err := token.NewToken().ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if data.Data.Type != model.TypeUserAdmin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "API only for admin"})
		c.Abort()
		return
	}

	c.Next()
}

func GetTokenFromGinContext(c *gin.Context) string {
	authorizationHeader := c.GetHeader("Authorization")

	authorizationValues := strings.SplitN(authorizationHeader, " ", 2)

	if len(authorizationValues) < 2 || strings.ToLower(authorizationValues[0]) != "bearer" {
		return ""
	}

	return strings.TrimSpace(authorizationValues[1])
}

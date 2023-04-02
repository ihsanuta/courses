package handler

import (
	"courses/app/model"
	"log"
	"net/http"
	"strconv"

	"courses/helper/cloudinary"

	"github.com/gin-gonic/gin"
)

func (h *handler) InsertCourse(c *gin.Context) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	uploadUrl, err := cloudinary.ImageUploadHelper(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	price, err := strconv.Atoi(c.Request.FormValue("price"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	category, err := strconv.Atoi(c.Request.FormValue("category_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	form := model.PayloadCourse{
		Name:       c.Request.FormValue("name"),
		Price:      price,
		CategoryID: category,
		Image:      uploadUrl,
	}

	result, err := h.usecase.Course.Insert(form)
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

func (h *handler) UpdateCourse(c *gin.Context) {
	var form model.PayloadCourse
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	result, err := h.usecase.Course.Update(form, id)
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

func (h *handler) GetCourse(c *gin.Context) {
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

	result, err := h.usecase.Course.GetByID(id)
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

func (h *handler) GetListCourse(c *gin.Context) {
	var form model.ParamsCourse
	if err := c.ShouldBindQuery(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("%v", form)

	result, err := h.usecase.Course.GetList(form)
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

func (h *handler) GetStatistic(c *gin.Context) {
	courses, free, err := h.usecase.Course.Total()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users, err := h.usecase.User.Total()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data": model.Total{
			TotalUser:        users,
			TotalCourses:     courses,
			TotalCOursesFree: free,
		},
	})
}

func (h *handler) GetCategory(c *gin.Context) {
	category, err := h.usecase.Course.GetCategoryCourse()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    category,
	})
}

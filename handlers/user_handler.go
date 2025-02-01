package handlers

import (
	"let-go/database"
	"let-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Lấy danh sách người dùng
// @Description Lấy toàn bộ danh sách người dùng
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// @Summary Thêm người dùng mới
// @Description Tạo người dùng mới với tên và tuổi
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Người dùng mới"
// @Success 201 {object} models.User
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newUser)
	c.JSON(http.StatusCreated, newUser)
}

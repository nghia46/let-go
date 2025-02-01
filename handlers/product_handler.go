package handlers

import (
	"let-go/database"
	"let-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Lấy danh sách sản phẩm
// @Description Lấy toàn bộ danh sách sản phẩm
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// @Summary Thêm sản phẩm mới
// @Description Tạo sản phẩm mới với tên và giá
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Sản phẩm mới"
// @Success 201 {object} models.Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

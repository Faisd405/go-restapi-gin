package example

import (
	"encoding/json"
	"net/http"

	ExampleModel "github.com/faisd405/go-restapi-gin/src/app/example/model"
	database "github.com/faisd405/go-restapi-gin/src/config"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var examples []ExampleModel.Example

	database.DB.Find(&examples)
	c.JSON(http.StatusOK, gin.H{"examples": examples})

}

func Show(c *gin.Context) {
	var example ExampleModel.Example
	id := c.Param("id")

	if err := database.DB.First(&example, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"example": example})
}

func Create(c *gin.Context) {

	var example ExampleModel.Example

	if err := c.ShouldBindJSON(&example); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&example)
	c.JSON(http.StatusOK, gin.H{"example": example})
}

func Update(c *gin.Context) {
	var example ExampleModel.Example
	id := c.Param("id")

	if err := c.ShouldBindJSON(&example); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&example).Where("id = ?", id).Updates(&example).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate example"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(c *gin.Context) {

	var example ExampleModel.Example

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if database.DB.Delete(&example, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus example"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

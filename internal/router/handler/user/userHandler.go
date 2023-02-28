package userHandler

import (
	"log"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"GinProject/internal/model"

	"GinProject/internal/db"
)

var (
	ds *gorm.DB
)

func Init() {
	ds = db.GetDBSession()
}

func ListUser(c *gin.Context) {
	var users []model.ApiUsers

	result := ds.Model(&model.Users{}).Select("acct", "fullname", "created_at", "updated_at").Find(&users)

	if result.Error != nil {
		log.Fatal(result.Error)
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "null",
		"data":  users,
	})
}

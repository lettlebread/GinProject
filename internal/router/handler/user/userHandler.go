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
	var dbResult *gorm.DB

	fullName := c.Query("fullname")

	if fullName != "" {
		dbResult = ds.Model(&model.Users{}).Where("fullname = ?", fullName).Select("acct", "fullname", "created_at", "updated_at").Find(&users)
	} else {
		dbResult = ds.Model(&model.Users{}).Select("acct", "fullname", "created_at", "updated_at").Find(&users)
	}

	if dbResult.Error != nil {
		log.Fatal(dbResult.Error)
		c.Error(dbResult.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "null",
		"data":  users,
	})
}

func GetUser(c *gin.Context) {
	var user model.ApiUsers
	var dbResult *gorm.DB

	account := c.Param("account")

	dbResult = ds.Model(&model.Users{}).Where("acct = ?", account).Select("acct", "fullname", "created_at", "updated_at").First(&user)

	if dbResult.Error != nil {
		log.Fatal(dbResult.Error)
		c.Error(dbResult.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "null",
		"data":  user,
	})
}

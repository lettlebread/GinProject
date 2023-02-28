package userHandler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"GinProject/internal/model"

	"GinProject/internal/db"
	bcryptUtil "GinProject/internal/util/bcrypt"

	validator "gopkg.in/validator.v2"
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
		fmt.Print("in error")
		log.Fatal(dbResult.Error)
		c.Error(dbResult.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "null",
		"data":  user,
	})
}

func CreateUser(c *gin.Context) {
	var dbResult *gorm.DB
	var payload model.CreateUserData
	err := c.Bind(&payload)
	if err != nil {
		log.Fatal(err)
		c.Error(err)
		return
	}

	err = validator.Validate(payload)
	if err != nil {
		log.Fatal(err)
		c.Error(err)
		return
	}

	hashedPwd, err := bcryptUtil.EncryptString(payload.Password)
	if err != nil {
		log.Fatal(err)
		c.Error(err)
		return
	}

	newUser := &model.Users{
		Acct:       payload.Account,
		Pwd:        hashedPwd,
		Fullname:   payload.Fullname,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	dbResult = ds.Create(&newUser)

	if dbResult.Error != nil {
		log.Fatal(dbResult.Error)
		c.Error(dbResult.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "null",
	})
}

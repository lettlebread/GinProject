package userHandler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"GinProject/internal/model"

	"GinProject/internal/db"
	bcryptUtil "GinProject/internal/util/bcrypt"
	jwtUtil "GinProject/internal/util/jwt"

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
	page, _ := strconv.Atoi(c.Query("page"))
	pagesize, _ := strconv.Atoi(c.Query("pagesize"))
	sortby := c.Query("sortby")

	dbResult = ds.Model(&model.Users{}).Select("acct", "fullname", "created_at", "updated_at")

	if fullName != "" {
		dbResult = dbResult.Where("fullname = ?", fullName)
	} else if page > 0 && pagesize > 0 {
		offset := (page - 1) * pagesize
		dbResult = dbResult.Offset(offset).Limit(pagesize)
	}

	if sortby != "" && isValidCol(sortby) {
		dbResult = dbResult.Order(sortby)
	}

	dbResult = dbResult.Find(&users)

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

func isValidCol(col string) bool {
	arr := []string{"acct", "fullname", "created_at", "updated_at"}

	for _, v := range arr {
		if v == col {
			return true
		}
	}
	return false
}

func GetUser(c *gin.Context) {
	var user model.ApiUsers
	var dbResult *gorm.DB

	account := c.Param("account")

	dbResult = ds.Model(&model.Users{}).Where("acct = ?", account).Select("acct", "fullname", "created_at", "updated_at").First(&user)

	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"error": "user not found",
			})
		} else {
			log.Fatal(dbResult.Error)
			c.Error(dbResult.Error)
		}

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

func Login(c *gin.Context) {
	var dbResult *gorm.DB
	var payload model.LoginUserData
	var user model.Users

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

	dbResult = ds.Model(&model.Users{}).Where("acct = ?", payload.Account).Select("acct", "pwd").First(&user)
	if dbResult.Error != nil {
		log.Fatal(dbResult.Error)
		c.Error(dbResult.Error)
		return
	}

	err = bcryptUtil.ComparePassword(user.Pwd, payload.Password)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "login failed",
		})
		return
	}

	token, err := jwtUtil.CreateToken(payload.Account)
	if err != nil {
		log.Fatal(err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "null",
		"token": token,
	})
}

func DeleteUser(c *gin.Context) {
	var dbResult *gorm.DB

	account := c.Param("account")
	currentUser := c.GetString("CURRENT_USER")

	if account == currentUser {
		c.JSON(http.StatusOK, gin.H{
			"error": "can't delete current user",
		})
		return
	}

	dbResult = ds.Where("acct = ?", account).Delete(&model.Users{})

	if dbResult.Error != nil {
		log.Fatal(dbResult.Error)
		c.Error(dbResult.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "null",
	})
}

func UpdateUser(c *gin.Context) {
	var dbResult *gorm.DB
	var payload model.UpdateUserData
	var user model.Users
	account := c.Param("account")

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

	dbResult = ds.Model(&model.Users{}).Where("acct = ?", account).Select("acct").First(&user)

	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"error": "user not found",
			})
		} else {
			log.Fatal(dbResult.Error)
			c.Error(dbResult.Error)
		}

		return
	}

	dbResult = ds.Model(&model.Users{}).Where("acct = ?", account).Updates(map[string]interface{}{"fullname": payload.Fullname, "updated_at": time.Now()})

	if dbResult.Error != nil {
		log.Fatal(dbResult.Error)
		c.Error(dbResult.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "null",
	})
}

func NoRoute(c *gin.Context) {
	c.JSON(404, gin.H{"error": "page not found"})
}

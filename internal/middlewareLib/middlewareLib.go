package middlewareLib

import (
	jwtUtil "GinProject/internal/util/jwt"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTValidate(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	authError := errors.New("invalid auth token")

	if authHeader == "" {
		c.AbortWithError(401, authError)
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.AbortWithError(401, authError)
	}

	account, err := jwtUtil.VarifyToken(parts[1])
	if err != nil {
		c.AbortWithError(401, authError)
	}

	c.Set("CURRENT_USER", account)
	c.Next()
}

func ErrorWrapper(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(500, gin.H{
			"error": c.Errors[0].Error(),
		})
	}
}

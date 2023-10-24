package internal

import (
	"time"

	"github.com/gin-gonic/gin"
)

// User struct definition
type User struct {
	ID             uint
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	CountryCode    string
	PasswordHash   string
	PasswordExpiry time.Time
}

type UserStorage interface {
	CreateUser(User) error
	GetUser(string) User
}

func CreateUser(c *gin.Context) {
	c.JSON(200, map[string]interface{}{"message": "user has been added"})
}

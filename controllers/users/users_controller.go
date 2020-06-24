package users

import (
	"fmt"
	"net/http"

	"github.com/c-petrini/bookstore_users-api/domain/users"
	"github.com/c-petrini/bookstore_users-api/services"
	"github.com/c-petrini/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.CreateUser(user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

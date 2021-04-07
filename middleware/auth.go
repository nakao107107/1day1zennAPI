package middleware

import (
	"1day1zennAPI/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Middleware struct {

}

func (m Middleware) CheckToken(c *gin.Context){
	var us service.UserService
	//headerにtokenがセットされていなければreturn
	if(c.GetHeader("Authorization") == ""){
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	_, err := us.GetUserByToken(c.GetHeader("Authorization"))
	if(err != nil){
		fmt.Println("there are any user")
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
}
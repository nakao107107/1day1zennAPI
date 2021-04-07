package controller

import (
	utils "1day1zennAPI/lib"
	"1day1zennAPI/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	_ "1day1zennAPI/lib"
)

type User struct {
	Login string `json:"login"`
}

//global
var ls service.LoginService
var us service.UserService

// ログイン GET /login
func (pc Controller) GetRedirectUrl(c *gin.Context) {
	loginUrl := ls.GetLoginUrl()
	fmt.Println(loginUrl)
	c.JSON(http.StatusOK, loginUrl)
}

func (pc Controller) GetToken(c *gin.Context){

	token, _ := ls.GetToken(c.Query("code"))
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := new(http.Client)
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	var u User
	_ = json.Unmarshal(byteArray, &u)

	//token生成
	accessToken := utils.RandomString(10)

	us.FindOrCreateUser(u.Login, accessToken)

	c.JSON(http.StatusOK, accessToken)
}


package controller

import (
	"1day1zennAPI/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var hs service.HistoryService

// 検索 GET /histories
func (pc Controller) Index(c *gin.Context) {
	user, _ := us.GetUserByToken(c.GetHeader("Authorization"))
	// 検索処理
	p, err := hs.GetHistories(user.Id)
	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

func (pc Controller) Create(c *gin.Context){
	user, _ := us.GetUserByToken(c.GetHeader("Authorization"))
	// 検索処理
	type BodyType struct {
		Type string `json:"type"`
		Url string `json:"url"`
		Impression string `json:"impression"`
	}

	var bodyType BodyType
	c.BindJSON(&bodyType)
	p, err := hs.CreateHistory(user.Id, bodyType.Type, bodyType.Url, bodyType.Impression)
	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, p)
	}
}
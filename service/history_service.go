package service

import (
	"1day1zennAPI/db"
	"1day1zennAPI/entity"
)

type history entity.History

type HistoryService struct {

}

// 履歴全権取得
func (s HistoryService) GetHistories(userId uint) ([]history, error) {

	// DB接続
	db := db.GetDB()

	var histories []history

	// DB接続確認
	if err := db.Take(&histories).Error; err != nil {
		return nil, err
	}

	// 履歴全件取得クエリ
	tx := db
	tx = tx.Where("user_id=?", userId).Find(&histories)

	return histories, nil
}

// 履歴作成
func (s HistoryService) CreateHistory(userId uint, siteType string, url string, impression string) (history, error) {

	// DB接続
	db := db.GetDB()

	history := history{UserId: userId, Type: siteType, Url: url, Impression: impression}

	db.Create(&history)

	return history, nil
}
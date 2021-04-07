package service

import (
	"1day1zennAPI/db"
	"1day1zennAPI/entity"
)

type user entity.User

type UserService struct {

}

// 履歴全件取得
func (us UserService) FindOrCreateUser(name string, token string) (user, error) {

	// DB接続
	db := db.GetDB()

	var user user

	_ = db.Take(&user)

	// 履歴全件取得クエリ

	if err := db.Where("github_id = ?", name).First(&user).Error; err != nil {
		// userが存在しなければ新規作成
		newUser := entity.User{}
		newUser.GithubId = name
		newUser.Token = token
		db.Create(&newUser)
	} else {
		user := entity.User{}
		user.GithubId = name
		user.Token = token
		db.Update(&user)
	}

	return user, nil
}

//tokenからUserを取得
func (us UserService) GetUserByToken(token string) (user entity.User, err error){
	// DB接続
	db := db.GetDB()

	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
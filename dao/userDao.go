package dao

import (
	"github.com/bubbly-hb/blogSystem-gin/db"
	"github.com/bubbly-hb/blogSystem-gin/model"
)

func GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Debug().Where("email=?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func GetUserByID(id uint) (*model.User, error) {
	user := &model.User{}
	err := db.DB.Debug().Where("id=?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func IsEmailExist(email string) bool {
	user := &model.User{}
	db.DB.Where("email=?", email).First(user)
	return user.ID != 0
}

func CreateUser(user *model.User) error {
	return db.DB.Create(user).Error
}

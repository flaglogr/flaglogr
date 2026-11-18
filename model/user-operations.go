package model

import (
	"context"

	"gorm.io/gorm"
)

func (db *DBConnector) GetUserByOIDCUserID(userid string) (User, bool) {
	user, err := gorm.G[User](db.DB).Where("uid = ?", userid).First(context.Background())
	if err != nil {
		return User{
			OidcId: userid,
		}, false
	}
	return user, true
}

func (db *DBConnector) UpdateOrCreateUser(userid string, username string, email string) uint {
	user, exists := db.GetUserByOIDCUserID(userid)
	wasUpdated := false
	if !exists {
		user = User{
			OidcId:   userid,
			Username: username,
		}
		wasUpdated = true
	} else {
		if user.Username != username {
			user.Username = username
			wasUpdated = true
		}
	}
	if wasUpdated {
		u := User{}
		db.DB.Save(&user).First(&u)
		return u.ID
	} else {
		return user.ID
	}
}

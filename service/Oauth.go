package service

import (
	"KNBot/model"
	"xorm.io/xorm"
)

type OauthService struct {
	db *xorm.Engine
}

//保存验证过的用户信息
func (os *OauthService) SaveOauthUser(BindUser model.User) error {
	_, err := os.db.Insert(&BindUser)
	return err
}

package service

import (
	"xorm.io/xorm"
)

type OauthService struct {
	db *xorm.Engine
}

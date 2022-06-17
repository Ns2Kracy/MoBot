package service

import (
	"xorm.io/xorm"
)

type UserService struct {
	Db *xorm.Engine
}

package OsuUser

import "xorm.io/xorm"

type OsuService struct {
	Db *xorm.Engine
}

func (os *OsuService) GetMode(state string) string {
	_, err := os.Db.Where(" qq = ? ", state).Cols("main_mode").Get("main_mode")
	if err != nil {
		return "0"
	}
	return ""
}

func (os *OsuService) UpdateMode(state string, mode int) bool {
	_, err := os.Db.Where(" qq = ? ", state).Update(&mode)
	if err != nil {
		return mode == 0
	}
	return mode != 0
}

package util

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 问候语句
var (
	GoodNightMessage     = "Good night!"
	GoodMorningMessage   = "Good morning!"
	GoodAfternoonMessage = "Good afternoon!"
	GoodEveningMessage   = "Good evening!"
)

package utils

import "github.com/gin-gonic/gin"

type Error struct {
	Code    string
	Message string
}

func SendOK(c *gin.Context, data gin.H) {
	SendDataSingle(c, 200, data)
}

func SendDataSingle(c *gin.Context, code int, data gin.H) {
	c.JSON(code, gin.H{"code": code, "data": data})
}

func SendDataArray(c *gin.Context, code int, array []gin.H) {
	c.JSON(code, gin.H{"code": code, "data": array})
}

func SendError(c *gin.Context, code int, errorObj Error) {
	c.JSON(code, gin.H{"code": code, "errcode": errorObj.Code, "err": errorObj.Message})
}

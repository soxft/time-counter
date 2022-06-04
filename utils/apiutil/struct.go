package apiutil

import "github.com/gin-gonic/gin"

type Apier interface {
	Out(success bool, msg string, data interface{})
	Success(msg string, data interface{})
	Fail(msg string)
	FailWithMsg(msg string, data interface{})
	Abort(httpCode int, msg string, errorCode int)
	Abort401(msg string, errCode int)
}

type Api struct {
	Ctx *gin.Context
}

package apiutil

import "github.com/gin-gonic/gin"

func New(ctx *gin.Context) *Api {
	return &Api{
		Ctx: ctx,
	}
}

func (c *Api) Out(success bool, msg string, data interface{}) {
	c.Ctx.JSON(200, gin.H{
		"success": success,
		"message": msg,
		"data":    data,
	})
}

func (c *Api) Success(msg string) {
	c.Out(true, msg, gin.H{})
}

func (c *Api) SuccessWithData(msg string, data interface{}) {
	c.Out(true, msg, data)
}

func (c *Api) Fail(msg string) {
	c.Out(false, msg, gin.H{})
}

func (c *Api) FailWithData(msg string, data interface{}) {
	c.Out(false, msg, data)
}

func (c *Api) Abort(httpCode int, msg string, errorCode int) {
	c.Ctx.AbortWithStatusJSON(httpCode, gin.H{
		"success": false,
		"message": msg,
		"data": gin.H{
			"errorCode": errorCode,
		},
	})
}

func (c *Api) Abort401(msg string, errorCode int) {
	c.Abort(401, msg, errorCode)
}

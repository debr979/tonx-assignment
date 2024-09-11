package controllers

import (
	"github.com/gin-gonic/gin"
)

type baseControllers struct {
	ctx *gin.Context
}

var Base baseControllers

func (r *baseControllers) Request(ctx *gin.Context) *baseControllers {
	r.ctx = ctx
	return r
}

func (r *baseControllers) ParseBody(req any) error {
	return r.ctx.ShouldBindJSON(req)
}

func (r *baseControllers) Response(httpCode int, errText string, data any) {
	var responseData = struct {
		ErrCode int    `json:"err_code"`
		Msg     string `json:"msg"`
		Data    any    `json:"data"`
	}{
		ErrCode: httpCode,
		Msg:     errText,
		Data:    data,
	}
	r.ctx.JSON(httpCode, responseData)
	return
}

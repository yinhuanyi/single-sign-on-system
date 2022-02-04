package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"user-server/controllers/response"
	"user-server/model"
	"user-server/utils"
)

// SSOLoginHandler 前端登录请求
func SSOLoginHandler(c *gin.Context)  {
	fmt.Println("c..................")
	response.ResponseSuccess(c, nil)
}

func GoodsCreateHandler(c *gin.Context)  {

	GoodsCreateInput := new(model.GoodsCreateInput)
	if err := c.ShouldBindJSON(GoodsCreateInput); err != nil {
		zap.L().Error("c.ShouldBindJSON", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, response.CodeInvalidParam)
			return
		}
		fmt.Printf("%v",utils.RemoveTopStruct(errs.Translate(utils.Trans)))
		response.ResponseErrorWithMsg(c, response.CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(utils.Trans)))
		return
	}

}



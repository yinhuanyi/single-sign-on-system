package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"test-servers/goods/controllers/response"
	"test-servers/goods/model"
	"test-servers/goods/utils"
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

func GoodsGetHandler(c *gin.Context)  {
	fmt.Println(c)
	fmt.Printf("%+v\n", c)

}

func GoodsUpdateHandler(c *gin.Context)  {

}

func GoodsDeleteHandler(c *gin.Context)  {

}

func GoodsListHandler(c *gin.Context)  {

	//categoryList, err := service.CategoryList()
	//if err != nil {
	//	zap.L().Error("service.CategoryList", zap.Error(err))
	//	response.ResponseError(c, response.CodeServerBusy)
	//	return
	//}
	//
	//response.ResponseSuccess(c, categoryList)

}


package libraryController

import (
	"funnel/app/controller"
	"funnel/app/errors"
	"funnel/app/service/libraryService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 图书馆历史借书记录
// @Description 图书馆借书记录（暂时只支持10本）
// @Tags 图书馆
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":[{...}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/libraryService/history/0 [post]
func LibraryBorrowHistory(context *gin.Context) {

	user, err := controller.LoginHandle(context, libraryService.GetUser)
	if err != nil {
		return
	}
	books, err := libraryService.GetBorrowHistory(user)

	if err == errors.ERR_Session_Expired {
		user, err = controller.LoginHandle(context, libraryService.GetUser)
		books, err = libraryService.GetBorrowHistory(user)
	}

	if err != nil {
		controller.ErrorHandle(context, err)
		return
	}
	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(books))
}

// @Summary 图书馆当前借书记录
// @Description 图书馆当前借书记录
// @Tags 图书馆
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":[{...}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/libraryService/current [post]
func LibraryCurrentBorrow(context *gin.Context) {
	user, err := controller.LoginHandle(context, libraryService.GetUser)
	if err != nil {
		return
	}
	books, err := libraryService.GetCurrentBorrow(user)

	if err == errors.ERR_Session_Expired {
		user, err = controller.LoginHandle(context, libraryService.GetUser)
		books, err = libraryService.GetCurrentBorrow(user)
	}

	if err != nil {
		controller.ErrorHandle(context, err)
		return
	}
	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(books))
}

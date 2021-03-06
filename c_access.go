package main

import (
	"github.com/kataras/iris"
	"net/http"
)

/**
* @api {post} /admin/login 用户登陆
* @apiName 用户登陆
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 用户登陆
* @apiSampleRequest /admin/login
* @apiParam {string} username 用户名
* @apiParam {string} password 密码
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CUserLogin(ctx iris.Context) {
	aul := new(UserJson)

	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(apiResource(false, nil, "请求参数错误"))
	} else {
		if UserNameErr := validate.Var(aul.Username, "required,min=4,max=20"); UserNameErr != nil {
			ctx.StatusCode(iris.StatusOK)
			ctx.JSON(apiResource(false, nil, "用户名格式错误"))
		} else if PwdErr := validate.Var(aul.Password, "required,min=5,max=20"); PwdErr != nil {
			ctx.StatusCode(iris.StatusOK)
			ctx.JSON(apiResource(false, nil, "密码格式错误"))
		} else {
			ctx.StatusCode(iris.StatusOK)
			response, status, msg := LUserAdminCheckLogin(aul.Username, aul.Password)
			ctx.JSON(apiResource(status, response, msg))
		}
	}
}

/**
* @api {get} /logout 用户退出登陆
* @apiName 用户退出登陆
* @apiGroup Users
* @apiVersion 1.0.0
* @apiDescription 用户退出登陆
* @apiSampleRequest /logout
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CUserLogout(ctx iris.Context) {

	aui := ctx.Values().GetString("auth_user_id")
	uid := uint(t.ParseInt(aui, 0))
	LUserAdminLogout(uid)

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(apiResource(true, nil, "退出登陆"))
}

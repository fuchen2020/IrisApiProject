package main

import (
	"github.com/kataras/iris"
)

/**
* @api {get} /perms 获取所有的权限
* @apiName 获取所有的权限
* @apiGroup Perms
* @apiVersion 1.0.0
* @apiDescription 获取所有的权限
* @apiSampleRequest /perms
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CGetAllPerms(ctx iris.Context) {
	offset := t.ParseInt(ctx.FormValue("offset"), 1)
	limit := t.ParseInt(ctx.FormValue("limit"), 20)
	name := ctx.FormValue("name")
	orderBy := ctx.FormValue("orderBy")

	perms := MGetAllPerms(name, orderBy, offset, limit)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(apiResource(true, perms, "操作成功"))
}

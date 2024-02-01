// 登录日志 控制层

package controller

import (
	"admin-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 分页获取登录日志列表
// @Summary 分页获取登录日志列表接口
// @Produce json
// @Description 分页获取登录日志列表接口
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param username query string false "用户名"
// @Param loginStatus query string false "登录状态（1-成功 2-失败）"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysLoginInfo/list [get]
// @Security ApiKeyAuth
func GetSysLoginInfoList(c *gin.Context) {
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	Username := c.Query("username")
	LoginStatus := c.Query("loginStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysLoginInfoService().GetSysLoginInfoList(c, Username, LoginStatus, BeginTime, EndTime, PageSize, PageNum)
}

// // 批量删除登录日志
// // @Summary 批量删除登录日志接口
// // @Produce json
// // @Description 批量删除登录日志接口
// // @Param data body entity.DelSysLoginInfoDto true "data"
// // @Success 200 {object} result.Result
// // @router /api/sysLoginInfo/batch/delete [delete]
// // @Security ApiKeyAuth
// func BatchDeleteSysLoginInfo(c *gin.Context) {
// 	var dto entity.DelSysLoginInfoDto
// 	_ = c.BindJSON(&dto)
// 	service.SysLoginInfoService().BatchDeleteSysLoginInfo(c, dto)
// }

// // 根据ID删除登录日志
// // @Summary 根据ID删除登录日志接口
// // @Produce json
// // @Description 根据ID删除登录日志接口
// // @Param data body entity.SysLoginInfoIdDto true "data"
// // @Success 200 {object} result.Result
// // @router /api/sysLoginInfo/delete [delete]
// // @Security ApiKeyAuth
// func DeleteSysLoginInfoById(c *gin.Context) {
// 	var dto entity.SysLoginInfoIdDto
// 	_ = c.BindJSON(&dto)
// 	service.SysLoginInfoService().DeleteSysLoginInfo(c, dto)
// }

// // 清空登录日志
// // @Summary 清空登录日志接口
// // @Produce json
// // @Description 清空登录日志接口
// // @Success 200 {object} result.Result
// // @router /api/sysLoginInfo/clean [delete]
// // @Security ApiKeyAuth
// func CleanSysLoginInfo(c *gin.Context) {
// 	service.SysLoginInfoService().CleanSysLoginInfo(c)
// }

// 岗位 控制层
// author xiaoRui

package controller

import (
	"admin-api/api/entity"
	"admin-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysPost entity.SysPost

// @Summary 新增岗位接口
// @Produce json
// @Description 新增岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/add [post]
// @Security ApiKeyAuth
func CreateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().CreateSysPost(c, sysPost)
}

// // 分页查询岗位列表
// // @Summary 分页查询岗位列表
// // @Produce json
// // @Description 分页查询岗位列表
// // @Param pageNum query int false "分页数"
// // @Param pageSize query int false "每页数"
// // @Param postName query string false "岗位名称"
// // @Param postStatus query string false "状态：1->启用,2->禁用"
// // @Param beginTime query string false "开始时间"
// // @Param endTime query string false "结束时间"
// // @Success 200 {object} result.Result
// // @router /api/post/list [get]
// // @Security ApiKeyAuth
func GetSysPostList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PostName := c.Query("postName")
	PostStatus := c.Query("postStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysPostService().GetSysPostList(c, PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
}

// // 根据id查询岗位
// // @Summary 根据id查询岗位
// // @Produce json
// // @Description 根据id查询岗位
// // @Param id query int true "ID"
// // @Success 200 {object} result.Result
// // @router /api/post/info [get]
// // @Security ApiKeyAuth
// func GetSysPostById(c *gin.Context) {
// 	Id, _ := strconv.Atoi(c.Query("id"))
// 	service.SysPostService().GetSysPostById(c, Id)
// }

// // 修改岗位
// // @Summary 修改岗位接口
// // @Produce json
// // @Description 修改岗位接口
// // @Param data body entity.SysPost true "data"
// // @Success 200 {object} result.Result
// // @router /api/post/update [put]
// // @Security ApiKeyAuth
// func UpdateSysPost(c *gin.Context) {
// 	_ = c.BindJSON(&sysPost)
// 	service.SysPostService().UpdateSysPost(c, sysPost)
// }

// // 根据id删除岗位
// // @Summary 根据id删除岗位接口
// // @Produce json
// // @Description 根据id删除岗位接口
// // @Param data body entity.SysPostIdDto true "data"
// // @Success 200 {object} result.Result
// // @router /api/post/delete [delete]
// // @Security ApiKeyAuth
// func DeleteSysPostById(c *gin.Context) {
// 	var dto entity.SysPostIdDto
// 	_ = c.BindJSON(&dto)
// 	service.SysPostService().DeleteSysPostById(c, dto)
// }

// // 批量删除岗位
// // @Summary 批量删除岗位接口
// // @Produce json
// // @Description 批量删除岗位接口
// // @Param data body entity.DelSysPostDto true "data"
// // @Success 200 {object} result.Result
// // @router /api/post/batch/delete [delete]
// // @Security ApiKeyAuth
// func BatchDeleteSysPost(c *gin.Context) {
// 	var dto entity.DelSysPostDto
// 	_ = c.BindJSON(&dto)
// 	service.SysPostService().BatchDeleteSysPost(c, dto)
// }

// // 岗位状态修改
// // @Summary 岗位状态修改接口
// // @Produce json
// // @Description 岗位状态修改接口
// // @Param data body entity.UpdateSysPostStatusDto true "data"
// // @Success 200 {object} result.Result
// // @router /api/post/updateStatus [put]
// // @Security ApiKeyAuth
// func UpdateSysPostStatus(c *gin.Context) {
// 	var dto entity.UpdateSysPostStatusDto
// 	_ = c.BindJSON(&dto)
// 	service.SysPostService().UpdateSysPostStatus(c, dto)
// }

// 岗位下拉列表
// @Summary 岗位下拉列表
// @Produce json
// @Description 岗位下拉列表
// @Success 200 {object} result.Result
// @router /api/post/vo/list [get]
// @Security ApiKeyAuth
func QuerySysPostVoList(c *gin.Context) {
	service.SysPostService().QuerySysPostVoList(c)
}

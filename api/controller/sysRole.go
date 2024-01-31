package controller

import (
	"admin-api/api/service"

	"github.com/gin-gonic/gin"
)

// 分页查询角色列表
// @Summary 分页查询角色列表接口
// @Produce json
// @Description 分页查询角色列表接口
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param roleName query string false "角色名称"
// @Param status query string false "帐号启用状态：1->启用,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/role/list [get]
// @Security ApiKeyAuth
func GetSysRoleList(c *gin.Context) {
	// RoleName := c.Query("roleName")
	// RoleStatus := c.Query("roleStatus")
	// service.SysRoleService().GetSysRoleList(c, RoleName, RoleStatus)
}

// 角色下拉列表
// @Summary 角色下拉列表
// @Produce json
// @Description 角色下拉列表
// @Success 200 {object} result.Result
// @router /api/role/vo/list [get]
// @Security ApikeyAuth
func QuerySysRoleVoList(c *gin.Context) {
	service.SysRoleService().QuerySysRoleVoList(c)
}

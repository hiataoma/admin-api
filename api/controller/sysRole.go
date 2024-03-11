package controller

import (
	"admin-api/api/entity"
	"admin-api/api/service"
	"strconv"

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
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	RoleName := c.Query("roleName")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysRoleService().GetSysRoleList(c, PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
}

// 根据id删除角色
// @Summary 根据id删除角色接口
// @Produce json
// @Description 根据id删除角色接口
// @Param data body entity.SysRoleIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/delete [delete]
// @Security ApikeyAuth
func DeleteSysRoleById(c *gin.Context) {
	var dto entity.SysRoleIdDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().DeleteSysRoleById(c, dto)
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

// 新增角色
// @Summary 新增角色接口
// @Produce json
// @Description 新增角色接口
// @Param data body entity.AddSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/add [post]
// @Security ApikeyAuth
func CreateSysRole(c *gin.Context) {
	var dto entity.AddSysRoleDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().CreateSysRole(c, dto)
}

// 角色详情
// @Summary 角色详情接口
// @Produce json
// @Description 角色详情接口
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/role/info/list [get]
// @Security ApikeyAuth
func GetSysRoleById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().GetSysRoleById(c, Id)
}

// 修改角色
// @Summary 修改角色
// @Produce json
// @Description 修改角色
// @Param data body entity.UpdateSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/update [put]
// @Security ApikeyAuth
func UpdateSysRole(c *gin.Context) {
	var dto entity.UpdateSysRoleDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().UpdateSysRole(c, dto)
}

// @Summary 修改角色转台
// @Produce json
// @Description 修改角色状态
// @Param data body entity.UpdateSysRoleStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/updateStatus [put]
// @Security ApikeyAuth
func UpdateSysRoleStatus(c *gin.Context) {
	var dto entity.UpdateSysRoleStatusDto
	_ = c.BindJSON(&dto)
	service.SysRoleService().UpdateSysRoleStatus(c, dto)
}

// 根据角色id查询菜单数据
// @Summary 根据角色id查询菜单数据接口
// @Produce json
// @Description 根据角色id查询菜单数据接口
// @Param id query int true "Id"
// @Success 200 {object} result.Result
// @router /api/role/vo/idList [get]
// @Security ApikeyAuth

func QueryRoleMenuIdList(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().QueryRoleMenuIdList(c, Id)
}

// AssignPermissions 分配权限
// @Summary 分配权限接口
// @Produce json
// @Description 分配权限接口
// @Param data body entity.RoleMenu true "data"
// @Success 200 {object} result.Result
// @router /api/role/assignPermissions [put]
// @Security ApiKeyAuth
func AssignPermissions(c *gin.Context) {
	var RoleMenu entity.RoleMenu
	_ = c.BindJSON(&RoleMenu)
	service.SysRoleService().AssignPermissions(c, RoleMenu)
}

// 菜单 控制层

package controller

import (
	"admin-api/api/entity"
	"admin-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sysMenu entity.SysMenu

// // 新增菜单
// // @Summary 新增菜单接口
// // @Produce json
// // @Description 新增菜单接口
// // @Param data body entity.SysMenu true "data"
// // @Success 200 {object} result.Result
// // @router /api/menu/add [post]
// // @Security ApiKeyAuth
// func CreateSysMenu(c *gin.Context) {
// 	_ = c.BindJSON(&sysMenu)
// 	service.SysMenuService().CreateSysMenu(c, sysMenu)
// }

// // 查询新增选项列表
// // @Summary 查询新增选项列表接口
// // @Produce json
// // @Description 查询新增选项列表接口
// // @Success 200 {object} result.Result
// // @router /api/menu/vo/list [get]
// // @Security ApiKeyAuth
func QuerySysMenuVoList(c *gin.Context) {
	service.SysMenuService().QuerySysMenuVoList(c)
}

// 根据id查询菜单
// @Summary 根据id查询菜单
// @Produce json
// @Description 根据id查询菜单
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/menu/info [get]
// @Security ApiKeyAuth
func GetSysMenu(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysMenuService().GetSysMenu(c, Id)
}

// 修改菜单
// @Summary 修改菜单接口
// @Produce json
// @Description 修改菜单接口
// @Param data body entity.SysMenu true "data"
// @Success 200 {object} result.Result
// @router /api/menu/update [put]
// @Security ApiKeyAuth
func UpdateSysMenu(c *gin.Context) {
	_ = c.BindJSON(&sysMenu)
	service.SysMenuService().UpdateSysMenu(c, sysMenu)
}

// // 根据id删除菜单
// // @Summary 根据id删除菜单接口
// // @Produce json
// // @Description 根据id删除菜单接口
// // @Param data body entity.SysMenuIdDto true "data"
// // @Success 200 {object} result.Result
// // @router /api/menu/delete [delete]
// // @Security ApiKeyAuth
// func DeleteSysMenu(c *gin.Context) {
// 	var dto entity.SysMenuIdDto
// 	_ = c.BindJSON(&dto)
// 	service.SysMenuService().DeleteSysMenu(c, dto)
// }

// 查询菜单列表
// @Summary 查询菜单列表
// @Produce json
// @Description 查询菜单列表
// @Param menuName query string false "菜单名称"
// @Param menuStatus query string false "菜单状态"
// @Success 200 {object} result.Result
// @router /api/menu/list [get]
// @Security ApiKeyAuth
func GetSysMenuList(c *gin.Context) {
	MenuName := c.Query("menuName")
	MenuStatus := c.Query("menuStatus")
	service.SysMenuService().GetSysMenuList(c, MenuName, MenuStatus)
}

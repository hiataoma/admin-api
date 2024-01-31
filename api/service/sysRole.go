package service

import (
	"admin-api/api/dao"
	"admin-api/common/result"

	"github.com/gin-gonic/gin"
)

type ISysRoleService interface {
	QuerySysRoleVoList(c *gin.Context)
}

type SysRoleServiceImpl struct{}

// 角色下拉列表
func (s SysRoleServiceImpl) QuerySysRoleVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysRoleVoList())
}

var sysRoleService = SysRoleServiceImpl{}

func SysRoleService() ISysRoleService {
	return &sysRoleService
}

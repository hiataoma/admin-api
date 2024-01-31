package dao

import (
	"admin-api/api/entity"
	. "admin-api/pkg/db"
)

// 角色下拉列表
func QuerySysRoleVoList() (sysRoleVo []entity.SysRoleVo) {
	Db.Table("sys_role").Select("id, role_name").Scan(&sysRoleVo)
	return sysRoleVo
}

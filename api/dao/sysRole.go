package dao

import (
	"admin-api/api/entity"
	"admin-api/common/util"
	. "admin-api/pkg/db"
	"time"
)

// 根据名称查询
func GetSysRoleByName(roleName string) (sysRole entity.SysRole) {
	Db.Where("role_name = ?", roleName).First(&sysRole)
	return sysRole
}

// 根据key查询
func GetSysRoleByKey(roleKey string) (sysRole entity.SysRole) {
	Db.Where("role_key = ?", roleKey).First(&sysRole)
	return sysRole
}

// 角色下拉列表
func QuerySysRoleVoList() (sysRoleVo []entity.SysRoleVo) {
	Db.Table("sys_role").Select("id, role_name").Scan(&sysRoleVo)
	return sysRoleVo
}

// 分页查询角色列表
func GetSysRoleList(PageNum, PageSize int, RoleName, status, BeginTime, EndTime string) (sysRole []*entity.SysRole, count int64) {
	curDb := Db.Table("sys_role")
	if RoleName != "" {
		curDb = curDb.Where("role_name = ?", RoleName)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if status != "" {
		curDb = curDb.Where("status = ?", status)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time DESC").Find(&sysRole)
	return sysRole, count
}

// 新增角色
func CreateSysRole(dto entity.AddSysRoleDto) bool {
	sysRoleByName := GetSysRoleByName(dto.RoleName)
	if sysRoleByName.ID > 0 {
		return false
	}
	sysRoleByKey := GetSysRoleByKey(dto.RoleKey)
	if sysRoleByKey.ID > 0 {
		return false
	}
	addSysRole := entity.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Description: dto.Description,
		Status:      dto.Status,
		CreateTime:  util.HTime{Time: time.Now()},
	}
	tx := Db.Create(&addSysRole)
	// 下面代码含义：如果影响的行数大于0，说明插入成功，返回true，否则返回false
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// 获取用户信息详情
func GetSysRoleById(Id int) (sysRole entity.SysRole) {
	Db.First(&sysRole, Id)
	return sysRole
}

// 更新角色
func UpdateSysRole(dto entity.UpdateSysRoleDto) (sysRole entity.SysRole) {
	Db.First(&sysRole, dto.Id)
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	sysRole.Status = dto.Status
	if dto.Description != "" {
		sysRole.Description = dto.Description
	}
	Db.Save(&sysRole)
	return sysRole
}

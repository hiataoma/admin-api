package dao

import (
	"admin-api/api/entity"
	"admin-api/common/util"
	. "admin-api/pkg/db"
	"time"
)

// 用户详情
func SysAdminDetail(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	username := dto.Username
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// 根据用户名查询用户
func GetSysAdminByUsername(username string) (sysAdmin entity.SysAdmin) {
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// 新增用户
func CreateSysAdmin(dto entity.AddSysAdminDto) bool {
	sysAdminByUsername := GetSysAdminByUsername(dto.Username)
	if sysAdminByUsername.ID > 0 {
		return false // 当前用户已经存在
	}
	sysAdmin := entity.SysAdmin{
		PostId:     dto.PostId,
		DeptId:     dto.DeptId,
		Username:   dto.Username,
		Nickname:   dto.Nickname,
		Password:   util.EncryptionMd5(dto.Password),
		Phone:      dto.Phone,
		Email:      dto.Email,
		Note:       dto.Note,
		Status:     dto.Status,
		CreateTime: util.HTime{Time: time.Now()},
	}
	tx := Db.Create(&sysAdmin)
	sysAdminExist := GetSysAdminByUsername(dto.Username)
	var entity entity.SysAdminRole
	entity.AdminId = sysAdminExist.ID
	entity.RoleId = dto.RoleId
	Db.Create(&entity)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// 根据id查询用户详情
func GetSysAdminInfo(Id int) (sysAdminInfo entity.SysAdminInfo) {
	Db.Table("sys_admin").
		Select("sys_admin.*, sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_admin_role.role_id = sys_role.id").
		First(&sysAdminInfo, Id)
	return sysAdminInfo
}

// 修改用户
func UpdateSysAdmin(dto entity.UpdateSysAdminDto) (sysAdmin entity.SysAdmin) {
	Db.First(&sysAdmin, dto.Id)
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	sysAdmin.PostId = dto.PostId
	sysAdmin.DeptId = dto.DeptId
	sysAdmin.Status = dto.Status
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	if dto.Note != "" {
		sysAdmin.Note = dto.Note
	}
	Db.Save(&sysAdmin)
	// 删除之前的角色，再分配新的角色 更新关联表里面的内容
	// 为什么不能直接更新关联表里面的内容呢？因为gorm的关联表更新是先删除再新增
	var sysAdminRole entity.SysAdminRole
	Db.Where("admin_id = ?", dto.Id).Delete(&entity.SysAdminRole{})
	sysAdminRole.AdminId = dto.Id
	sysAdminRole.RoleId = dto.RoleId
	Db.Create(&sysAdminRole)
	return sysAdmin
}

// // 根据id删除用户
// func DeleteSysAdminById(dto entity.SysAdminIdDto) {
// 	Db.First(&entity.SysAdmin{}, dto.Id)
// 	Db.Delete(&entity.SysAdmin{}, dto.Id)
// 	Db.Where("admin_id = ?", dto.Id).Delete(&entity.SysAdminRole{})
// }

// // 修改用户状态
// func UpdateSysAdminStatus(dto entity.UpdateSysAdminStatusDto) {
// 	var sysAdmin entity.SysAdmin
// 	Db.First(&sysAdmin, dto.Id)
// 	sysAdmin.Status = dto.Status
// 	Db.Save(&sysAdmin)
// }

// 重置密码
func ResetSysAdminPassword(dto entity.ResetSysAdminPasswordDto) {
	var sysAdmin entity.SysAdmin
	Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = util.EncryptionMd5(dto.Password)
	Db.Save(&sysAdmin)
}

// 分页查询用户列表
func GetSysAdminList(PageSize, PageNum int, Username, Status, BeginTime, EndTime string) (sysAdminVo []entity.SysAdminVo, count int64) {
	curDb := Db.Table("sys_admin").
		Select("sys_admin.*, sys_post.post_name, sys_role.role_name, sys_dept.dept_name").
		Joins("LEFT JOIN sys_post ON sys_admin.post_id = sys_post.id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_role.id = sys_admin_role.role_id").
		Joins("LEFT JOIN sys_dept ON sys_dept.id = sys_admin.dept_id")
	if Username != "" {
		curDb = curDb.Where("sys_admin.username = ?", Username)
	}
	if Status != "" {
		curDb = curDb.Where("sys_admin.status = ?", Status)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("sys_admin.create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("sys_admin.create_time DESC").Find(&sysAdminVo)
	return sysAdminVo, count
}

// // 修改个人信息
// func UpdatePersonal(dto entity.UpdatePersonalDto) (sysAdmin entity.SysAdmin) {
// 	Db.First(&sysAdmin, dto.Id)
// 	if dto.Icon != "" {
// 		sysAdmin.Icon = dto.Icon
// 	}
// 	if dto.Username != "" {
// 		sysAdmin.Username = dto.Username
// 	}
// 	if dto.Nickname != "" {
// 		sysAdmin.Nickname = dto.Nickname
// 	}
// 	if dto.Phone != "" {
// 		sysAdmin.Phone = dto.Phone
// 	}
// 	if dto.Email != "" {
// 		sysAdmin.Email = dto.Email
// 	}
// 	Db.Save(&sysAdmin)
// 	return sysAdmin
// }

// // 修改个人密码
// func UpdatePersonalPassword(dto entity.UpdatePersonalPasswordDto) (sysAdmin entity.SysAdmin) {
// 	Db.First(&sysAdmin, dto.Id)
// 	sysAdmin.Password = dto.NewPassword
// 	Db.Save(&sysAdmin)
// 	return sysAdmin
// }

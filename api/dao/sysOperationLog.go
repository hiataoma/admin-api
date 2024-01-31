package dao

import (
	"admin-api/api/entity"
	. "admin-api/pkg/db"
)

// 新增操作日志
func CreateSysOperationLog(log entity.SysOperationLog) {
	Db.Create(&log)
}

// 分页查询操作日志列表
func GetSysOperationLogList(Username, BeginTime, EndTime string, PageSize, PageNum int) (sysOperationLog []entity.SysOperationLog, count int64) {
	curDb := Db.Table("sys_operation_log")
	if Username != "" {
		curDb = curDb.Where("username =?", Username)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("`create_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&sysOperationLog)
	return sysOperationLog, count
}

// 根据id删除操作日志
func DeleteSysOperationLogById(dto entity.SysOperationLogIdDto) {
	Db.Delete(&entity.SysOperationLog{}, dto)
}

// 批量删除批量操作日志
func BatchDeleteSysOperationLog(dto entity.BatchDeleteSysOperationLogDto) {
	Db.Where("id in (?)", dto.Ids).Delete(&entity.SysOperationLog{})
}

// 清空操作日志
func CleanSysOperationLog() {
	Db.Exec("truncate table sys_operation_log")
}

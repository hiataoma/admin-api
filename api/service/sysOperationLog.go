// 操作日志 数据层

package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"

	"github.com/gin-gonic/gin"
)

type ISysOperationLogService interface {
	GetSysOperationLogList(c *gin.Context, Username, BeginTime, EndTime string, PageSize, PageNum int)
	DeleteSysOperationLogById(c *gin.Context, dto entity.SysOperationLogIdDto)
	BatchDeleteSysOperationLog(c *gin.Context, dto entity.BatchDeleteSysOperationLogDto)
	CleanSysOperationLog(c *gin.Context)
}

type SysOperationLogServiceImpl struct{}

// 清空操作日志
func (s SysOperationLogServiceImpl) CleanSysOperationLog(c *gin.Context) {
	dao.CleanSysOperationLog()
	result.Success(c, true)
}

// 批量删除操作日志
func (s SysOperationLogServiceImpl) BatchDeleteSysOperationLog(c *gin.Context, dto entity.BatchDeleteSysOperationLogDto) {
	dao.BatchDeleteSysOperationLog(dto)
	result.Success(c, true)
}

// 根据id删除操作日志
func (s SysOperationLogServiceImpl) DeleteSysOperationLogById(c *gin.Context, dto entity.SysOperationLogIdDto) {
	dao.DeleteSysOperationLogById(dto)
	result.Success(c, true)
}

// 分页查询操作日志列表
func (s SysOperationLogServiceImpl) GetSysOperationLogList(c *gin.Context, Username, BeginTime, EndTime string, PageSize, PageNum int) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysOperationLog, count := dao.GetSysOperationLogList(Username, BeginTime, EndTime, PageSize, PageNum)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysOperationLog})
}

var sysOperationLogService = SysOperationLogServiceImpl{}

func SysOperationLogService() ISysOperationLogService {
	return &sysOperationLogService
}

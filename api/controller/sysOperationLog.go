package controller

import (
	"admin-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 分页获取操作日志列表
// @Summary 分页获取操作日志列表接口
// @Produce json
// @Description 分页获取操作日志列表接口
// @Param pageSize query int false "每页数"
// @Param pageNum query int false "分页数"
// @Param username query string false "用户名"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/list [get]
// @Security ApiKeyAuth
func GetSysOperationLogList(c *gin.Context) {
	Username := c.Query("username")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	service.SysOperationLogService().GetSysOperationLogList(c, Username, BeginTime, EndTime, PageSize, PageNum)
}

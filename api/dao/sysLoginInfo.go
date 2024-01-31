package dao

import (
	"admin-api/api/entity"
	"admin-api/common/util"
	"fmt"
	"time"
)

// 新增登录日志
func CreateSysLoginInfo(username, ipAddress, loginLocation, browser, os, message string, loginStatus int) {
	sysLoginInfo := entity.SysLoginInfo{
		Username:      username,
		IpAddress:     ipAddress,
		LoginLocation: loginLocation,
		Browser:       browser,
		Os:            os,
		Message:       message,
		LoginStatus:   loginStatus,
		LoginTime:     util.HTime{Time: time.Now()},
	}
	fmt.Print(sysLoginInfo)
	// Db.Save(&sysLoginInfo)
}

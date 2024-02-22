package controller

import (
	"github.com/gin-gonic/gin"
)

// @Summary welcome接口
// @Produce json
// @Description welcome接口
// @Param data body entity.LoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/login [post]
func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 10000,
		"data": gin.H{
			"name": "mahaitao",
		},
		"msg": "成功",
	})
}

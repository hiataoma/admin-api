// 岗位 服务层

package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"

	"github.com/gin-gonic/gin"
)

type ISysPostService interface {
	CreateSysPost(c *gin.Context, sysPost entity.SysPost)
	GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string)
	GetSysPostById(c *gin.Context, Id int)
	UpdateSysPost(c *gin.Context, sysPost entity.SysPost)
	DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto)
	BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto)
	UpdateSysPostStatus(c *gin.Context, dto entity.UpdateSysPostStatusDto)
	QuerySysPostVoList(c *gin.Context)
}

// 接口
type SysPostServiceImpl struct{}

// 岗位下拉列表
func (s SysPostServiceImpl) QuerySysPostVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysPostVoList())
}

// 修改状态
func (s SysPostServiceImpl) UpdateSysPostStatus(c *gin.Context, dto entity.UpdateSysPostStatusDto) {
	dao.UpdateSysPostStatus(dto)
	result.Success(c, true)
}

// 批量删除
func (s SysPostServiceImpl) BatchDeleteSysPost(c *gin.Context, dto entity.DelSysPostDto) {
	dao.BatchDeleteSysPost(dto)
	result.Success(c, true)
}

// 根据id删除岗位
func (s SysPostServiceImpl) DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto) {
	dao.DeleteSysPostById(dto)
	result.Success(c, true)
}

// 修改岗位
func (s SysPostServiceImpl) UpdateSysPost(c *gin.Context, sysPost entity.SysPost) {
	result.Success(c, dao.UpdateSysPost(sysPost))
}

// 根据id查询岗位
func (s SysPostServiceImpl) GetSysPostById(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysPostById(Id))
}

// 分页查询岗位列表
func (s SysPostServiceImpl) GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysPost, count := dao.GetSysPostList(PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysPost})
}

// 新增岗位
func (s SysPostServiceImpl) CreateSysPost(c *gin.Context, sysPost entity.SysPost) {
	bool := dao.CreateSysPost(sysPost)
	if !bool {
		result.Failed(c, int(result.ApiCode.POSTALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.POSTALREADYEXISTS))
		return
	}
	result.Success(c, true)
}

var sysPostService = SysPostServiceImpl{}

func SysPostService() ISysPostService {
	return &sysPostService
}

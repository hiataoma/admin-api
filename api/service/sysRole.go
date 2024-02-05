package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"

	"github.com/gin-gonic/gin"
)

type ISysRoleService interface {
	QuerySysRoleVoList(c *gin.Context)
	GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, Status, BeginTime, EndTime string)
	CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto)
	GetSysRoleById(c *gin.Context, Id int)
	UpdateSysRole(c *gin.Context, dto entity.UpdateSysRoleDto)
	UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto)
	QueryRoleMenuIdList(c *gin.Context, Id int)
	AssignPermissions(c *gin.Context, menu entity.RoleMenu)
}

type SysRoleServiceImpl struct{}

// 角色下拉列表
func (s SysRoleServiceImpl) QuerySysRoleVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysRoleVoList())
}

// 角色分页列表
func (s SysRoleServiceImpl) GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysRole, count := dao.GetSysRoleList(PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysRole})
	return
}

// 分配权限
func (s *SysRoleServiceImpl) AssignPermissions(c *gin.Context, menu entity.RoleMenu) {
	result.Success(c, dao.AssignPermissions(menu))
}

// 新增角色

func (s SysRoleServiceImpl) CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto) {
	bool := dao.CreateSysRole(dto)
	if !bool {
		result.Failed(c, int(result.ApiCode.ROLENAMEALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.ROLENAMEALREADYEXISTS))
		return
	}
	result.Success(c, true)
}

// 更新角色
func (s SysRoleServiceImpl) UpdateSysRole(c *gin.Context, dto entity.UpdateSysRoleDto) {
	sysRole := dao.UpdateSysRole(dto)
	result.Success(c, sysRole)
}

// 获取用户信息详情
func (s SysRoleServiceImpl) GetSysRoleById(c *gin.Context, Id int) {
	sysRole := dao.GetSysRoleById(Id)
	result.Success(c, sysRole)
}

// 根据角色id查询菜单数据
func (s *SysRoleServiceImpl) QueryRoleMenuIdList(c *gin.Context, Id int) {
	roleMenuIdList := dao.QueryRoleMenuIdList(Id)
	var idList = make([]int, 0)
	for _, id := range roleMenuIdList {
		idList = append(idList, id.Id)
	}
	result.Success(c, idList)
}

// 更新角色状态
func (s SysRoleServiceImpl) UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto) {
	result.Success(c, dao.UpdateSysRoleStatus(dto))
}

var sysRoleService = SysRoleServiceImpl{}

func SysRoleService() ISysRoleService {
	return &sysRoleService
}

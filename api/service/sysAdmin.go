package service

import (
	"admin-api/api/dao"
	"admin-api/api/entity"
	"admin-api/common/result"
	"admin-api/common/util"
	"admin-api/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// 定义接口
type ISysAminService interface {
	Login(c *gin.Context, dto entity.LoginDto)
	CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto)
	GetSysAdminInfo(c *gin.Context, Id int)
	UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto)
	// DeleteSysAdminById(c *gin.Context, dto entity.SysAdminIdDto)
	// UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto)
	ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto)
	GetSysAdminList(c *gin.Context, PageSize, PageNum int, Username, Status, BeginTime, EndTime string)
	// UpdatePersonal(c *gin.Context, dto entity.UpdatePersonalDto)
	// UpdatePersonalPassword(c *gin.Context, dto entity.UpdatePersonalPasswordDto)
}

type SysAdminServiceImpl struct{}

// 用户登录
func (s SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParameter), result.ApiCode.GetMessage(result.ApiCode.MissingLoginParameter))
		return
	}
	ip := c.ClientIP()
	// 验证码是否过期
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		dao.CreateSysLoginInfo(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "验证码已过期", 2)
		result.Failed(c, int(result.ApiCode.VerificationCodeHasExpired), result.ApiCode.GetMessage(result.ApiCode.VerificationCodeHasExpired))
		return
	}
	// 校验验证码
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		dao.CreateSysLoginInfo(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "验证码不正确", 2)
		result.Failed(c, int(result.ApiCode.CAPTCHANOTTRUE), result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}
	// 校验
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
		dao.CreateSysLoginInfo(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "密码不正确", 2)
		result.Failed(c, int(result.ApiCode.PASSWORDNOTTRUE), result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}
	const status int = 2
	if sysAdmin.Status == status {
		dao.CreateSysLoginInfo(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "账号已停用", 2)
		result.Failed(c, int(result.ApiCode.STATUSISENABLE), result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}
	// 生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(sysAdmin)
	dao.CreateSysLoginInfo(dto.Username, ip, util.GetRealAddressByIP(ip), util.GetBrowser(c), util.GetOs(c), "登录成功", 1)
	// 左侧菜单列表
	var leftMenuVo []entity.LeftMenuVo
	leftMenuList := dao.QueryLeftMenuList(sysAdmin.ID)
	for _, value := range leftMenuList {
		menuSvoList := dao.QueryMenuVoList(sysAdmin.ID, value.Id)
		item := entity.LeftMenuVo{}
		item.MenuSvoList = menuSvoList
		item.Id = value.Id
		item.MenuName = value.MenuName
		item.Icon = value.Icon
		item.Url = value.Url
		leftMenuVo = append(leftMenuVo, item)
	}
	// 权限列表
	permissionList := dao.QueryPermissionList(sysAdmin.ID)
	var stringList = make([]string, 0)
	for _, value := range permissionList {
		stringList = append(stringList, value.Value)
	}
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdmin, "leftMenuList": leftMenuVo, "permissionList": stringList})
}

// 新增用户
func (s SysAdminServiceImpl) CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingNewAdminParameter), result.ApiCode.GetMessage(result.ApiCode.MissingNewAdminParameter))
		return
	}
	bool := dao.CreateSysAdmin(dto)
	if !bool {
		result.Failed(c, int(result.ApiCode.USERNAMEALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.USERNAMEALREADYEXISTS))
		return
	}
	result.Success(c, bool)
	return
}

// 用户列表
func (s SysAdminServiceImpl) GetSysAdminList(c *gin.Context, PageSize, PageNum int, Username, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysAdmin, count := dao.GetSysAdminList(PageSize, PageNum, Username, Status, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysAdmin})
	return
}

// 用户详情
func (s SysAdminServiceImpl) GetSysAdminInfo(c *gin.Context, Id int) {
	result.Success(c, dao.GetSysAdminInfo(Id))
}

// 重置密码
func (s SysAdminServiceImpl) ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto) {
	dao.ResetSysAdminPassword(dto)
	result.Success(c, true)
}

// 用户更新
func (s SysAdminServiceImpl) UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto) {
	result.Success(c, dao.UpdateSysAdmin(dto))
}

var sysAdminService = SysAdminServiceImpl{}

func SysAdminService() ISysAminService {
	return &sysAdminService
}

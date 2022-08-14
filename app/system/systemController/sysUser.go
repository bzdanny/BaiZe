package systemController

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/app/utils"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
)

type UserController struct {
	us systemService.IUserService
	ps systemService.IPostService
	rs systemService.IRoleService
}

func NewUserController(us *systemServiceImpl.UserService, ps *systemServiceImpl.PostService, rs *systemServiceImpl.RoleService) *UserController {
	return &UserController{us: us, ps: ps, rs: rs}
}

// ChangeStatus 修改用户状态
// @Summary 修改用户状态
// @Description 修改用户状态
// @Tags 用户相关
// @Param  object body systemModels.EditUserStatus true "用户信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData  "成功"
// @Router /system/user/changeStatus [put]
func (uc *UserController) ChangeStatus(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)

	sysUser := new(systemModels.EditUserStatus)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		bzc.ParameterError()
		return
	}
	sysUser.SetUpdateBy(bzc.GetUserId())
	uc.us.UpdateUserStatus(sysUser)
	bzc.Success()
}

// ResetPwd 重置密码
// @Summary 重置密码
// @Description 重置密码
// @Tags 用户相关
// @Param  object body systemModels.ResetPwd true "密码"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData  "成功"
// @Router /system/user/resetPwd [put]
func (uc *UserController) ResetPwd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)

	resetPwd := new(systemModels.ResetPwd)
	if err := c.ShouldBindJSON(resetPwd); err != nil {
		bzc.ParameterError()
		return
	}
	uc.us.ResetPwd(resetPwd.UserId, resetPwd.Password)
	bzc.Success()

}

// UserEdit 修改用户
// @Summary 修改用户
// @Description 修改用户
// @Tags 用户相关
// @Param  object body systemModels.SysUserEdit true "用户信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData  "成功"
// @Router /system/user  [put]
func (uc *UserController) UserEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)

	sysUser := new(systemModels.SysUserEdit)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		fmt.Println(err)
		bzc.ParameterError()
		return
	}
	if uc.us.CheckPhoneUnique(sysUser.UserId, sysUser.Phonenumber) {
		bzc.Waring("新增用户'" + sysUser.Phonenumber + "'失败，手机号码已存在")
		return
	}
	if uc.us.CheckEmailUnique(sysUser.UserId, sysUser.Email) {
		bzc.Waring("新增用户'" + sysUser.Email + "'失败，邮箱账号已存在")
		return
	}
	sysUser.SetUpdateBy(bzc.GetUserId())
	uc.us.UpdateUser(sysUser)
	bzc.Success()
}

// UserAdd 添加用户
// @Summary 添加用户
// @Description 添加用户
// @Tags 用户相关
// @Param  object body systemModels.SysUserAdd true "用户信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData  "成功"
// @Router /system/user  [post]
func (uc *UserController) UserAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)

	user := bzc.GetUser()
	sysUser := new(systemModels.SysUserAdd)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		bzc.ParameterError()
		return
	}
	if sysUser.DeptId == nil {
		sysUser.DeptId = user.DeptId
	}
	if uc.us.CheckUserNameUnique(sysUser.UserName) {
		bzc.Waring("新增用户'" + sysUser.UserName + "'失败，登录账号已存在")
		return
	}
	if uc.us.CheckPhoneUnique(sysUser.UserId, sysUser.Phonenumber) {
		bzc.Waring("新增用户'" + sysUser.Phonenumber + "'失败，手机号码已存在")
		return
	}

	if uc.us.CheckEmailUnique(sysUser.UserId, sysUser.Email) {
		bzc.Waring("新增用户'" + sysUser.Email + "'失败，邮箱账号已存在")
		return
	}
	sysUser.SetCreateBy(user.UserId)
	uc.us.InsertUser(sysUser)
	bzc.Success()
}

// UserList 添加用户
// @Summary 添加用户
// @Description 添加用户
// @Tags 用户相关
// @Param  object query systemModels.SysUserDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData{data=commonModels.ListData{Rows=[]systemModels.SysUserVo}}  "成功"
// @Router /system/user  [get]
func (uc *UserController) UserList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	user := new(systemModels.SysUserDQL)
	_ = c.ShouldBind(user)
	user.SetDataScope(bzc.GetUser(), "d", "u")
	list, count := uc.us.SelectUserList(user)
	bzc.SuccessListData(list, count)

}

// UserGetInfo 获取当前用户信息
// @Summary 获取当前用户信息
// @Description 获取当前用户信息
// @Tags 用户相关
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData{data=systemModels.UserInfo}  "成功"
// @Router /system/user/  [get]
func (uc *UserController) UserGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	ui := new(systemModels.UserInfo)
	ui.Posts = uc.ps.SelectPostAll()
	user := bzc.GetUser()
	role := new(systemModels.SysRoleDQL)
	c.ShouldBind(role)
	role.SetDataScope(user, "d", "")
	roleList := uc.rs.SelectRoleAll(role)
	if !utils.IsAdmin(user.UserId) {
		for i, role := range roleList {
			if role.RoleId == 1 {
				roleList = append(roleList[:i], roleList[i+1:]...)
				break
			}
		}
	}
	ui.Roles = roleList
	bzc.SuccessData(ui)

}

// UserAuthRole 根据用户编号获取授权角色
// @Summary 根据用户编号获取授权角色
// @Description 根据用户编号获取授权角色
// @Tags 用户相关
// @Param id path string true "userId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData{data=systemModels.Auth}  "成功"
// @Router /system/user/authRole/{userId}  [get]
func (uc *UserController) UserAuthRole(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	userId := bzc.ParamInt64("userId")
	if userId == 0 {
		bzc.ParameterError()
	}
	ar := new(systemModels.Auth)
	ar.User = uc.us.SelectUserById(userId)
	role := new(systemModels.SysRoleDQL)
	user := bzc.GetUser()
	role.SetDataScope(user, "d", "")
	roles := uc.rs.SelectRoleAll(role)
	if !utils.IsAdmin(user.UserId) {
		for i, role := range roles {
			if role.RoleId == 1 {
				roles = append(roles[:i], roles[i+1:]...)
				break
			}
		}
	}
	ar.Roles = roles
	ar.RoleIds = gconv.Strings(uc.rs.SelectRoleListByUserId(userId))
	bzc.SuccessData(ar)
}

// UserGetInfoById 根据用户ID获取用户信息
// @Summary 根据用户ID获取用户信息
// @Description 根据用户ID获取用户信息
// @Tags 用户相关
// @Param id path string true "userId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  commonModels.ResponseData{data=systemModels.Auth}  "成功"
// @Router /system/user/{userId}  [get]
func (uc *UserController) UserGetInfoById(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	userId := bzc.ParamInt64("userId")
	if userId == 0 {
		bzc.ParameterError()
	}
	ar := new(systemModels.Auth)
	postList := uc.ps.SelectPostAll()

	user := bzc.GetUser()
	role := new(systemModels.SysRoleDQL)
	role.SetDataScope(user, "d", "")
	roleList := uc.rs.SelectRoleAll(role)
	if !utils.IsAdmin(user.UserId) {
		for i, role := range roleList {
			if role.RoleId == 1 {
				roleList = append(roleList[:i], roleList[i+1:]...)
				break
			}
		}
	}
	ar.Posts = postList
	ar.Roles = roleList
	ar.PostIds = gconv.Strings(uc.ps.SelectPostListByUserId(userId))
	ar.RoleIds = gconv.Strings(uc.rs.SelectRoleListByUserId(userId))
	ar.User = uc.us.SelectUserById(userId)
	bzc.SuccessData(ar)

}

// UserRemove 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags 系统用户
// @Param userIds path  []string true "userIds"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} commonModels.ResponseData
// @Router /system/user/{userIds} [delete]
func (uc *UserController) UserRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	uc.us.DeleteUserByIds(bzc.ParamInt64Array("userIds"))
	bzc.Success()
}

func (uc *UserController) UserImportData(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)

	fileHeader, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	file, _ := fileHeader.Open()
	defer file.Close()
	excelFile, _ := excelize.OpenReader(file)
	rows := excelFile.GetRows("Sheet1")
	loginUser := bzc.GetUser()
	data, num := uc.us.UserImportData(rows, loginUser.UserName, loginUser.DeptId)
	if num > 0 {
		bzc.Waring(data)
		return
	}
	bzc.SuccessMsg(data)
}

func (uc *UserController) UserExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	user := new(systemModels.SysUserDQL)
	_ = c.ShouldBind(user)
	user.SetDataScope(bzc.GetUser(), "d", "u")
	bzc.DataPackageExcel(uc.us.UserExport(user))
	return
}

func (uc *UserController) ImportTemplate(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	data := uc.us.ImportTemplate()
	bzc.DataPackageExcel(data)
	return
}

func (uc *UserController) InsertAuthRole(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	array := bzc.QueryInt64Array("roleIds")
	uc.us.InsertUserAuth(bzc.QueryInt64("userId"), array)
	bzc.Success()
	return
}

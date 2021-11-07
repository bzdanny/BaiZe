package systemController

import (
	"baize/app/system/service/loginService"
	"baize/app/system/service/loginService/loginServiceImpl"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
)

var iLogin loginService.ILoginService = loginServiceImpl.GetLoginService()
var iUser systemService.IUserService = systemServiceImpl.GetUserService()
var iMenu systemService.IMenuService = systemServiceImpl.GetMenuService()
var iDept systemService.IDeptService = systemServiceImpl.GetDeptService()
var iDictData systemService.IDictDataService = systemServiceImpl.GetDictDataService()
var iDictType systemService.IDictTypeService = systemServiceImpl.GetDictTypeService()
var iPost systemService.IPostService = systemServiceImpl.GetPostService()
var iRole systemService.IRoleService = systemServiceImpl.GetRoleService()

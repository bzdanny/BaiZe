package systemController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewTenantController, NewUserController, NewRoleController, NewPostController, NewPermissionController, NewLoginController, NewDictTypeController,
	NewProfileController, NewDictDataController, NewDeptController)

type SystemController struct {
	DictData   *DictDataController
	DictType   *DictTypeController
	Login      *LoginController
	User       *UserController
	Dept       *DeptController
	Role       *RoleController
	Post       *PostController
	Permission *PermissionController
	Profile    *ProfileController
}

func NewTenantController(
	DictData *DictDataController,
	DictType *DictTypeController,
	Login *LoginController,
	User *UserController,
	Dept *DeptController,
	Role *RoleController,
	Post *PostController,
	Permission *PermissionController,
	Profile *ProfileController,
) *SystemController {
	return &SystemController{
		DictData:   DictData,
		DictType:   DictType,
		Login:      Login,
		User:       User,
		Dept:       Dept,
		Role:       Role,
		Post:       Post,
		Permission: Permission,
		Profile:    Profile,
	}
}

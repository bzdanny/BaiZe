package systemController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewTenantController, NewUserController, NewRoleController, NewPostController, NewMenuController, NewLoginController, NewDictTypeController,
	NewProfileController, NewDictDataController, NewDeptController)

type SystemController struct {
	DictData *DictDataController
	DictType *DictTypeController
	Login    *LoginController
	User     *UserController
	Dept     *DeptController
	Role     *RoleController
	Post     *PostController
	Menu     *MenuController
	Profile  *ProfileController
}

func NewTenantController(
	DictData *DictDataController,
	DictType *DictTypeController,
	Login *LoginController,
	User *UserController,
	Dept *DeptController,
	Role *RoleController,
	Post *PostController,
	Menu *MenuController,
	Profile *ProfileController,
) *SystemController {
	return &SystemController{
		DictData: DictData,
		DictType: DictType,
		Login:    Login,
		User:     User,
		Dept:     Dept,
		Role:     Role,
		Post:     Post,
		Menu:     Menu,
		Profile:  Profile,
	}
}

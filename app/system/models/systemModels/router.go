package systemModels

type RouterVo struct {
	//路由名字
	Name string `json:"name"`
	//路由地址
	Path string `json:"path"`
	//是否隐藏路由，当设置 true 的时候该路由不会再侧边栏出现
	Hidden bool `json:"hidden"`
	//重定向地址，当设置 noRedirect 的时候该路由在面包屑导航中不可被点击
	Redirect string `json:"redirect"`

	Component string `json:"component"`

	AlwaysShow bool `json:"always_show"`

	Meta MetaVo `json:"meta"`

	Children []*RouterVo `json:"children"`
}

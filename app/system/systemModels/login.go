package systemModels

import (
	"github.com/bzdanny/BaiZe/baize/baizeEntity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Uuid     string `json:"uuid" binding:"required"`
}

type LoginUser struct {
	Token         string
	LoginTime     int64
	ExpireTime    int64
	IpAddr        string
	LoginLocation string
	Browser       string
	Os            string
	User          *User
	RolePerms     []string
	Permissions   []string
}

type User struct {
	UserId      int64                  `json:"userId,string" db:"user_id" :"user_id"`
	DeptId      *int64                 `json:"deptId" db:"dept_id" :"dept_id"`
	UserName    string                 `json:"userName" db:"user_name" :"user_name"`
	NickName    string                 `json:"nickName" db:"nick_name" :"nick_name"`
	Email       *string                `json:"email" db:"email" :"email"`
	Phonenumber *string                `json:"phonenumber"db:"phonenumber" :"phonenumber"`
	Sex         string                 `json:"sex" db:"sex" :"sex"`
	Avatar      *string                `json:"avatar" db:"avatar" :"avatar"`
	Password    string                 `json:"-" db:"password" :"password"`
	Status      string                 `json:"status" db:"status" :"status"`
	DelFlag     string                 `json:"delFlag" db:"del_flag" :"del_flag"`
	LoginIp     string                 `json:"loginIp" db:"login_ip" :"login_ip"`
	LoginDate   *time.Time             `json:"loginDate" db:"login_date" :"login_date"`
	Remark      *string                `json:"remark" db:"remark" :"remark"`
	ParentId    *int64                 `json:"parentId" db:"parent_id" :"parent_id"`
	DeptName    *string                `json:"deptName" db:"dept_name" :"dept_name"`
	CreateTime  *baizeEntity.BaiZeTime `json:"createTime" db:"create_time" :"create_time"`
	Roles       []*baizeEntity.Role    `json:"roles"`
}

func (u *User) GetRoles() []*baizeEntity.Role {

	return u.Roles
}

func (u *User) GetDeptId() int64 {

	if u == nil {
		return 0
	}
	return *u.DeptId
}

type JWT struct {
	TokenId string
	jwt.StandardClaims
}

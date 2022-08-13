package systemModels

import (
	"github.com/bzdanny/BaiZe/baize/baizeEntity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginBody struct {
	Username string `json:"username" binding:"required"` //用户名
	Password string `json:"password" binding:"required"` //密码
	Code     string `json:"code" binding:"required"`     //验证码
	Uuid     string `json:"uuid" binding:"required"`     //uuid
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
	UserId      int64                  `json:"userId,string" db:"user_id" db:"user_id"`
	DeptId      *int64                 `json:"deptId" db:"dept_id" db:"dept_id"`
	UserName    string                 `json:"userName" db:"user_name" db:"user_name"`
	NickName    string                 `json:"nickName" db:"nick_name" db:"nick_name"`
	Email       *string                `json:"email" db:"email" db:"email"`
	Phonenumber *string                `json:"phonenumber"db:"phonenumber" db:"phonenumber"`
	Sex         string                 `json:"sex" db:"sex" db:"sex"`
	Avatar      *string                `json:"avatar" db:"avatar" db:"avatar"`
	Password    string                 `json:"-" db:"password" db:"password"`
	Status      string                 `json:"status" db:"status" db:"status"`
	DelFlag     string                 `json:"delFlag" db:"del_flag" db:"del_flag"`
	LoginIp     string                 `json:"loginIp" db:"login_ip" db:"login_ip"`
	LoginDate   *time.Time             `json:"loginDate" db:"login_date" db:"login_date"`
	Remark      *string                `json:"remark" db:"remark" db:"remark"`
	ParentId    *int64                 `json:"parentId" db:"parent_id" db:"parent_id"`
	DeptName    *string                `json:"deptName" db:"dept_name" db:"dept_name"`
	CreateTime  *baizeEntity.BaiZeTime `json:"createTime" db:"create_time" db:"create_time"`
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

type GetInfo struct {
	User        *User    `json:"user"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

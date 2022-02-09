package loginModels

import (
	"baize/app/common/baize/baizeUnix"
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
	UserId      int64                `json:"userId,string" db:"user_id"`
	DeptId      *int64               `json:"deptId" db:"dept_id"`
	UserName    string               `json:"userName" db:"user_name"`
	NickName    string               `json:"nickName" db:"nick_name"`
	Email       *string              `json:"email" db:"email"`
	Phonenumber *string              `json:"phonenumber"db:"phonenumber"`
	Sex         string               `json:"sex" db:"sex"`
	Avatar      *string              `json:"avatar" db:"avatar"`
	Password    string               `json:"-" db:"password"`
	Status      string               `json:"status" db:"status"`
	DelFlag     string               `json:"delFlag" db:"del_flag"`
	LoginIp     string               `json:"loginIp" db:"login_ip"`
	LoginDate   *time.Time           `json:"loginDate" db:"login_date"`
	Remark      *string              `json:"remark" db:"remark"`
	ParentId    *int64               `json:"parentId" db:"parent_id"`
	DeptName    *string              `json:"deptName" db:"dept_name"`
	CreateTime  *baizeUnix.BaiZeTime `json:"createTime" db:"create_time"`
	Roles       []*Role
}

type Role struct {
	RoleId    int64  `db:"role_id"`
	DataScope string `db:"data_scope"`
}

type JWT struct {
	TokenId string
	jwt.StandardClaims
}

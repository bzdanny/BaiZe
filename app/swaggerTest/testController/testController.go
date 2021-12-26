package testController

import (
	"baize/app/common/commonModels"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type UserEntity struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

var users = make(map[int64]UserEntity)

func init() {
	users[1] = UserEntity{UserId: 1, UserName: "demo1", Password: "admin123", Mobile: "18688888888"}
	users[2] = UserEntity{UserId: 2, UserName: "demo2", Password: "admin123", Mobile: "18666666666"}
}

// DemoUserList 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags 演示用户相关
// @Produce application/json
// @Success 200 {object} commonModels.ResponseData
// @Router /test/user/list [get]
func DemoUserList(c *gin.Context) {
	entities := make([]UserEntity, 0, len(users))
	for _, user := range users {
		entities = append(entities, user)
	}
	i := int64(len(entities))
	c.JSON(http.StatusOK, commonModels.SuccessListData(entities, &i))
}

// GetUser 获取用户详细
// @Summary 获取用户详细
// @Description 获取用户详细
// @Tags 演示用户相关
// @Param userId path int true "用户ID"
// @Produce application/json
// @Success 200 {object} commonModels.ResponseData
// @Router /test/user/{userId} [get]
func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	entity := users[userId]
	if entity.UserId == 0 {
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	c.JSON(http.StatusOK, commonModels.SuccessData(entity))
}

// Save 新增用户
// @Summary 新增用户
// @Description 新增用户
// @Tags 演示用户相关
// @Param  object body testController.UserEntity true "用户信息"
// @Produce application/json
// @Success 200 {object} commonModels.ResponseData
// @Router /test/user [post]
func Save(c *gin.Context) {
	user := new(UserEntity)
	if err := c.ShouldBindJSON(user); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	fmt.Println(users[user.UserId])
	if user.UserId == 0 || users[user.UserId].UserId != 0 {
		c.JSON(http.StatusOK, commonModels.ErrorMsg("用户ID不能为空"))
		return
	}
	users[user.UserId] = *user
	c.JSON(http.StatusOK, commonModels.Success())
}

// Update 更新用户
// @Summary 更新用户
// @Description 更新用户
// @Tags 演示用户相关
// @Param  object body testController.UserEntity true "用户信息"
// @Produce application/json
// @Success 200 {object} commonModels.ResponseData
// @Router /test/user [put]
func Update(c *gin.Context) {
	user := new(UserEntity)
	if err := c.ShouldBindJSON(user); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	fmt.Println(users[user.UserId])
	if user.UserId == 0 || users[user.UserId].UserId == 0 {
		c.JSON(http.StatusOK, commonModels.ErrorMsg("用户不存在"))
		return
	}
	users[user.UserId] = *user
	c.JSON(http.StatusOK, commonModels.Success())
}

// Delete 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags 演示用户相关
// @Param userId path int true "用户ID"
// @Produce application/json
// @Success 200 {object} commonModels.ResponseData
// @Router /test/user/{userId} [delete]
func Delete(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	delete(users, userId)
	c.JSON(http.StatusOK, commonModels.Success())
}

package systemController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func PostList(c *gin.Context) {
	post := new(systemModels.SysPostDQL)
	c.ShouldBind(post)
	var page = commonModels.NewPageDomain()
	c.ShouldBind(post)
	post.SetLimit(page)

	list, count := iPost.SelectPostList(post)

	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func PostExport(c *gin.Context) {

}

func PostGetInfo(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("postId"), 10, 64)
	if err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	sysUser := iPost.SelectPostById(postId)
	c.JSON(http.StatusOK, commonModels.SuccessData(sysUser))
}

func PostAdd(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	sysPost := new(systemModels.SysPostDML)
	if err := c.ShouldBindJSON(sysPost); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		c.JSON(http.StatusOK, commonModels.ParameterError())
		return
	}
	sysPost.SetCreateBy(loginUser.User.UserName)
	iPost.InsertPost(sysPost)

	c.JSON(http.StatusOK, commonModels.Success())
}

func PostEdit(c *gin.Context) {
	loginUser := commonController.GetCurrentLoginUser(c)
	post := new(systemModels.SysPostDML)
	c.ShouldBindJSON(post)
	post.SetUpdateBy(loginUser.User.UserName)
	iPost.UpdatePost(post)
	c.JSON(http.StatusOK, commonModels.Success())

}

func PostRemove(c *gin.Context) {
	var s slicesUtils.Slices = strings.Split(c.Param("postIds"), ",")
	iPost.DeletePostByIds(s.StrSlicesToInt())
	c.JSON(http.StatusOK, commonModels.Success())
}

func Postoptionselect(c *gin.Context) {

}

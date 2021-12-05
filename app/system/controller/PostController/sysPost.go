package PostController

import (
	"baize/app/common/commonController"
	"baize/app/common/commonLog"
	"baize/app/common/commonModels"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

var iPost systemService.IPostService = systemServiceImpl.GetPostService()

func PostList(c *gin.Context) {
	post := new(systemModels.SysPostDQL)
	c.ShouldBind(post)
	post.SetLimit(c)
	list, count := iPost.SelectPostList(post)
	c.JSON(http.StatusOK, commonModels.SuccessListData(list, count))

}

func PostExport(c *gin.Context) {
	commonLog.SetLog(c, "岗位管理", "EXPORT")
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
	commonLog.SetLog(c, "岗位管理", "INSERT")
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
	commonLog.SetLog(c, "岗位管理", "UPDATE")
	loginUser := commonController.GetCurrentLoginUser(c)
	post := new(systemModels.SysPostDML)
	c.ShouldBindJSON(post)
	post.SetUpdateBy(loginUser.User.UserName)
	iPost.UpdatePost(post)
	c.JSON(http.StatusOK, commonModels.Success())

}

func PostRemove(c *gin.Context) {
	commonLog.SetLog(c, "岗位管理", "DELETE")
	var s slicesUtils.Slices = strings.Split(c.Param("postIds"), ",")
	iPost.DeletePostByIds(s.StrSlicesToInt())
	c.JSON(http.StatusOK, commonModels.Success())
}

package PostController

import (
	"baize/app/common/baize/baizeContext"
	"baize/app/system/models/systemModels"
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
	"baize/app/utils/slicesUtils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

var iPost systemService.IPostService = systemServiceImpl.GetPostService()

func PostList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	post := new(systemModels.SysPostDQL)
	c.ShouldBind(post)
	post.SetLimit(c)
	list, count := iPost.SelectPostList(post)
	bzc.SuccessListData(list, count)

}

func PostExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	post := new(systemModels.SysPostDQL)
	c.ShouldBind(post)
	data := iPost.PostExport(post)
	bzc.DataPackageExcel(data)
}

func PostGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	postId := bzc.ParamInt64("postId")
	if postId == 0 {
		zap.L().Error("参数错误")
		bzc.ParameterError()
		return
	}
	bzc.SuccessData(iPost.SelectPostById(postId))
}

func PostAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("岗位管理", "INSERT")
	sysPost := new(systemModels.SysPostDML)
	if err := c.ShouldBindJSON(sysPost); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		bzc.ParameterError()
		return
	}
	sysPost.SetCreateBy(bzc.GetCurrentUserName())
	iPost.InsertPost(sysPost)
	bzc.Success()
}

func PostEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("岗位管理", "UPDATE")
	post := new(systemModels.SysPostDML)
	c.ShouldBindJSON(post)
	post.SetUpdateBy(bzc.GetCurrentUserName())
	iPost.UpdatePost(post)
	bzc.Success()

}

func PostRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	bzc.SetLog("岗位管理", "DELETE")
	var s slicesUtils.Slices = strings.Split(c.Param("postIds"), ",")
	iPost.DeletePostByIds(s.StrSlicesToInt())
	bzc.Success()
}

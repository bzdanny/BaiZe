package systemController

import (
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/system/systemService"
	"github.com/bzdanny/BaiZe/app/system/systemService/systemServiceImpl"
	"github.com/bzdanny/BaiZe/baize/baizeContext"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	ps systemService.IPostService
}

func NewPostController(ps *systemServiceImpl.PostService) *PostController {
	return &PostController{ps: ps}
}

func (pc *PostController) PostList(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	post := new(systemModels.SysPostDQL)
	_ = c.ShouldBind(post)
	list, count := pc.ps.SelectPostList(post)
	bzc.SuccessListData(list, count)

}

func (pc *PostController) PostExport(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	post := new(systemModels.SysPostDQL)
	_ = c.ShouldBind(post)
	data := pc.ps.PostExport(post)
	bzc.DataPackageExcel(data)
}

func (pc *PostController) PostGetInfo(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	postId := bzc.ParamInt64("postId")
	if postId == 0 {
		bzc.ParameterError()
		return
	}
	bzc.SuccessData(pc.ps.SelectPostById(postId))
}

func (pc *PostController) PostAdd(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	sysPost := new(systemModels.SysPostAdd)
	if err := c.ShouldBindJSON(sysPost); err != nil {
		bzc.ParameterError()
		return
	}
	sysPost.SetCreateBy(bzc.GetUserId())
	pc.ps.InsertPost(sysPost)
	bzc.Success()
}

func (pc *PostController) PostEdit(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	post := new(systemModels.SysPostEdit)
	if err := c.ShouldBindJSON(post); err != nil {
		bzc.ParameterError()
		return
	}
	post.SetUpdateBy(bzc.GetUserId())
	pc.ps.UpdatePost(post)
	bzc.Success()

}

func (pc *PostController) PostRemove(c *gin.Context) {
	bzc := baizeContext.NewBaiZeContext(c)
	pc.ps.DeletePostByIds(bzc.ParamInt64Array("postIds"))
	bzc.Success()
}

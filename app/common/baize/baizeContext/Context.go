package baizeContext

import (
	"baize/app/common/commonModels"
	"baize/app/constant/business"
	"baize/app/constant/constants"
	"baize/app/constant/httpStatus"
	"baize/app/monitor/monitorModels"
	"baize/app/system/models/loginModels"
	"baize/app/utils/IpUtils"
	"baize/app/utils/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/mssola/user_agent"
	"net/http"
	"strconv"
	"strings"
)

type BaiZeContext struct {
	*gin.Context
}

func NewBaiZeContext(c *gin.Context) *BaiZeContext {
	return &BaiZeContext{c}
}

func (bzc *BaiZeContext) ParamInt64(key string) int64 {
	return gconv.Int64(bzc.Param(key))
}
func (bzc *BaiZeContext) ParamInt64Array(key string) []int64 {
	split := strings.Split(bzc.Param(key), ",")
	list := make([]int64, 0, len(split))
	for _, s := range split {
		list = append(list, gconv.Int64(s))
	}
	return list
}
func (bzc *BaiZeContext) QueryInt64(key string) int64 {
	return gconv.Int64(bzc.Query(key))
}
func (bzc *BaiZeContext) QueryInt64Array(key string) []int64 {
	split := strings.Split(bzc.Query(key), ",")
	list := make([]int64, 0, len(split))
	for _, s := range split {
		list = append(list, gconv.Int64(s))
	}
	return list
}

func (bzc *BaiZeContext) GetCurrentLoginUser() (loginUser *loginModels.LoginUser) {
	loginUserKey, _ := bzc.Get(constants.LoginUserKey)
	if loginUserKey != nil {
		loginUser = loginUserKey.(*loginModels.LoginUser)
	}
	return
}

func (bzc *BaiZeContext) SetLog(title string, businessTy business.BusinessType) {
	ol := new(monitorModels.SysOpenLog)
	ol.Title = title
	ol.BusinessType = businessTy.Msg()
	ol.Status = business.Success.Msg()
	ol.OperIp = bzc.ClientIP()
	ol.OperUrl = bzc.Request.URL.Path
	ol.RequestMethod = bzc.Request.Method
	loginUser := bzc.GetCurrentLoginUser()
	if loginUser != nil {
		ol.OperName = loginUser.User.UserName
	}
	bzc.Set(constants.LogKey, ol)
}
func (bzc *BaiZeContext) SetUserAgent(login *monitorModels.Logininfor) {
	login.InfoId = snowflake.GenID()
	ua := user_agent.New(bzc.Request.Header.Get("User-Agent"))
	ip := bzc.ClientIP()
	login.IpAddr = ip
	login.Os = ua.OS()
	login.LoginLocation = IpUtils.GetRealAddressByIP(ip)
	login.Browser, _ = ua.Browser()
}

func (bzc *BaiZeContext) Success() {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Success, Msg: httpStatus.Success.Msg()})
}
func (bzc *BaiZeContext) SuccessMsg(msg string) {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Success, Msg: msg})
}

func (bzc *BaiZeContext) SuccessData(data interface{}) {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Success, Msg: httpStatus.Success.Msg(), Data: data})
}
func (bzc *BaiZeContext) SuccessListData(rows interface{}, total *int64) {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Success, Msg: httpStatus.Success.Msg(), Data: commonModels.ListData{Rows: rows, Total: total}})
}

func (bzc *BaiZeContext) Waring(msg string) {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Waring, Msg: msg})
}
func (bzc *BaiZeContext) BzError() {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Error, Msg: httpStatus.Error.Msg()})
}
func (bzc *BaiZeContext) ErrorMsg(msg string) {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Error, Msg: msg})
}

func (bzc *BaiZeContext) ParameterError() {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Parameter, Msg: httpStatus.Parameter.Msg()})
}
func (bzc *BaiZeContext) InvalidToken() {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Unauthorized, Msg: httpStatus.Unauthorized.Msg()})
}
func (bzc *BaiZeContext) PermissionDenied() {
	bzc.JSON(http.StatusOK, commonModels.ResponseData{Code: httpStatus.Forbidden, Msg: httpStatus.Forbidden.Msg()})
}
func (bzc *BaiZeContext) DataPackageExcel(data []byte) {
	bzc.Header("Content-Type", "application/vnd.ms-excel")
	bzc.Header("Pragma", "public")
	bzc.Header("Cache-Control", "no-store")
	bzc.Header("Cache-Control", "max-age=0")
	bzc.Header("Content-Length", strconv.Itoa(len(data)))
	bzc.Data(http.StatusOK, "application/vnd.ms-excel", data)
}

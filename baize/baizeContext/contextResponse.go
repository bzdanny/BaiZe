package baizeContext

import (
	"github.com/bzdanny/BaiZe/app/commonModels"
	"github.com/bzdanny/BaiZe/app/constant/httpStatus"
	"net/http"
	"strconv"
)

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
func (bzc *BaiZeContext) DataPackageZip(data []byte) {
	bzc.Header("Access-Control-Allow-Origin", "*")
	bzc.Header("Access-Control-Expose-Headers", "Content-Disposition")
	bzc.Header("Content-Disposition", "attachment; filename=\"baize.zip\"")
	bzc.Header("Content-Length", strconv.Itoa(len(data)))
	bzc.Data(http.StatusOK, "application/octet-stream; charset=UTF-8", data)
}

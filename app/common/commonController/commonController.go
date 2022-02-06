package commonController

//func GetCurrentLoginUser(c *gin.Context) (loginUser *loginModels.LoginUser) {
//	loginUserKey, _ := c.Get(constants.LoginUserKey)
//	if loginUserKey != nil {
//		loginUser = loginUserKey.(*loginModels.LoginUser)
//	}
//	return
//}
//
//func DataPackageExcel (c *gin.Context,data []byte)  {
//	c.Header("Content-Type", "application/vnd.ms-excel")
//	c.Header("Pragma", "public")
//	c.Header("Cache-Control", "no-store")
//	c.Header("Cache-Control", "max-age=0")
//	c.Header("Content-Length", strconv.Itoa(len(data)))
//	c.Data(http.StatusOK, "application/vnd.ms-excel", data)
//}

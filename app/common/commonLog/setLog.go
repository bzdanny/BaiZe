package commonLog

//func SetLog(c *gin.Context, title string, businessTy business.BusinessType) {
//	ol := new(monitorModels.SysOpenLog)
//	ol.Title = title
//	ol.BusinessType = businessTy.Msg()
//	ol.Status = business.Success.Msg()
//	ol.OperIp = c.ClientIP()
//	ol.OperUrl = c.Request.URL.Path
//	ol.RequestMethod = c.Request.Method
//	loginUser := commonController.GetCurrentLoginUser(c)
//	if loginUser != nil {
//		ol.OperName = loginUser.User.UserName
//	}
//	c.Set(constants.LogKey, ol)
//}

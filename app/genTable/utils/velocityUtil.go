package genUtils

import (
	"baize/app/genTable/genTableModels"
	"baize/app/utils/stringUtils"
	"fmt"
	"strings"
)

func GetTemplateList() []string {
	templates := make([]string, 0, 11)
	templates = append(templates, "go/model.go.vm")
	templates = append(templates, "go/controller.go.vm")
	templates = append(templates, "go/iService.go.vm")
	templates = append(templates, "go/serviceImpl.go.vm")
	templates = append(templates, "go/iDao.go.vm")
	templates = append(templates, "go/daoImpl.go.vm")
	templates = append(templates, "go/router.go.vm")
	templates = append(templates, "sql/sql.sql.vm")
	templates = append(templates, "js/api.js.vm")
	templates = append(templates, "vue/index.vue.vm")
	return templates
}
func GetFileName(template string, genTable *genTableModels.GenTableVo) string {
	if strings.HasSuffix(template, "model.go.vm") {
		return fmt.Sprintf("/go/%s/%sModels/%s.go", genTable.ModuleName, genTable.PackageName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "controller.go.vm") {
		return fmt.Sprintf("/go/%s/%sController/%s.go", genTable.ModuleName, genTable.PackageName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "iService.go.vm") {
		return fmt.Sprintf("/go/%s/%sService/i%s.go", genTable.ModuleName, genTable.PackageName, stringUtils.Capitalize(genTable.BusinessName))
	}
	if strings.HasSuffix(template, "serviceImpl.go.vm") {
		return fmt.Sprintf("/go/%s/%sService/%sServiceImpl/%sImpl.go", genTable.ModuleName, genTable.PackageName, genTable.PackageName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "iDao.go.vm") {
		return fmt.Sprintf("/go/%s/%sDao/i%s.go", genTable.ModuleName, genTable.PackageName, stringUtils.Capitalize(genTable.BusinessName))
	}
	if strings.HasSuffix(template, "daoImpl.go.vm") {
		return fmt.Sprintf("/go/%s/%sDao/%sDaoImpl/%sImpl.go", genTable.ModuleName, genTable.PackageName, genTable.PackageName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "router.go.vm") {
		return fmt.Sprintf("/go/routes/%sRoutes/%sRouter.go", genTable.ModuleName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "sql.sql.vm") {
		return fmt.Sprintf("/go/%s.sql", genTable.BusinessName)
	}
	if strings.HasSuffix(template, "api.js.vm") {
		return fmt.Sprintf("/vue/api/%s/%s.js", genTable.ModuleName, genTable.BusinessName)
	}
	if strings.HasSuffix(template, "index.vue.vm") {
		return fmt.Sprintf("/vue/views/%s/%s/index.vue", genTable.ModuleName, genTable.BusinessName)
	}
	return ""
}

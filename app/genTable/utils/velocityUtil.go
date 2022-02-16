package genUtils

func GetTemplateList() []string {
	templates := make([]string, 0, 11)
	templates = append(templates, "go/controller/controller.go.vm")
	templates = append(templates, "go/dao/iDao.go.vm")
	templates = append(templates, "go/dao/daoImpl/daoImpl.go.vm")
	templates = append(templates, "go/model/model.go.vm")
	templates = append(templates, "go/routes/router.go.vm")
	templates = append(templates, "go/service/iService.go.vm")
	templates = append(templates, "go/service/serviceImpl/serviceImpl.go.vm")
	templates = append(templates, "js/api.js.vm")
	templates = append(templates, "sql/sql.sql.vm")
	templates = append(templates, "vue/index.vue.vm")
	return templates
}

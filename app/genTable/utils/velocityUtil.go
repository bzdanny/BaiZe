package genUtils

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

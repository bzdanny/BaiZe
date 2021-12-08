package task

var testMap = make(map[string]func())

func AddTask(invokeTarget string, run func()) {
	testMap[invokeTarget] = run
}
func GetByName(invokeTarget string) (run func()) {
	return testMap[invokeTarget]
}

func IsExistFunc(invokeTarget string)(ok bool){
	_,ok=testMap[invokeTarget]
	return
}
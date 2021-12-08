package task

import (
	"baize/app/quartz/task/taskUtils"
	"encoding/json"
	"fmt"
)

func init() {
	AddTask("NoParams", NoParams)
	AddTask("Params", Params)

}

func NoParams() {
	fmt.Println("无参测试")
}

type taskDemo struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func Params() {
	cache := taskUtils.GetQuartzCache("Params")
	t := new(taskDemo)
	fmt.Println(cache)
	err := json.Unmarshal([]byte(cache), t)
	if err != nil {
		fmt.Println("参数有误")
	}
	fmt.Print("有参数测试：")
	fmt.Println(t)
}

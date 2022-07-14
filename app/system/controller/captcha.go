package controller

import (
	"github.com/gin-gonic/gin"
)

type Test struct {
}

func NewTest() *Test {
	return new(Test)
}

func (t *Test) GetCode(c *gin.Context) {
	c.JSON(200, "Ok")
}

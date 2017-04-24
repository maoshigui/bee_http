package controllers

import (
	"github.com/astaxie/beego"
)

type TestcaseController struct {
	beego.Controller
}

func (c *TestcaseController) Get() {
	c.TplName = "testcase.html"
}

package controllers

import (
	"github.com/astaxie/beego"
)

type StatisticController struct {
	beego.Controller
}

func (c *StatisticController) Get() {
	c.TplName = "statistic.html"
}

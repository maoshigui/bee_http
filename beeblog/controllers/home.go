package controllers

import (
	"bee_http/beeblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "home.html"

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
}

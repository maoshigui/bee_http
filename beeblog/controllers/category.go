package controllers

import (
	"bee_http/beeblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")

	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("category", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 302)
		return
	}

	var err error
	c.Data["IsCategory"] = true
	c.Data["Categories"], err = models.GetAllCategories()
	c.TplName = "category.html"

	if err != nil {
		beego.Error(err)
	}
}

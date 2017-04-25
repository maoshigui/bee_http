package controllers

import (
	"bee_http/goms/models"

	"github.com/astaxie/beego"
)

type CellController struct {
	beego.Controller
}

func (c *CellController) Get() {
	switch c.Input().Get("op") {
	case "del":
		id := c.Input().Get("id")
		err := models.DelCell(id)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/cell", 302)
		return
	case "setup":
		c.Redirect("/cell", 302)
		return
	}

	var err error
	c.Data["IsCell"] = true
	c.Data["Cells"], err = models.GetAllCells()
	c.TplName = "cell.html"
	if err != nil {
		beego.Error(err)
	}
}

func (c *CellController) Post() {
	cuid := c.Input().Get("cuid")
	duid := c.Input().Get("duid")
	cellid := c.Input().Get("cellid")

	var err error
	err = models.AddCell(cuid, duid, cellid)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/cell", 302)
}

package controllers

import (
	"bee_http/goms/models"

	"github.com/astaxie/beego"
)

type CellController struct {
	beego.Controller
}

func (c *CellController) Get() {
	c.Data["IsCell"] = true
	c.TplName = "cell.html"

	cells, err := models.GetAllCells()
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Cells"] = cells
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

	c.Redirect("/topic", 302)
}

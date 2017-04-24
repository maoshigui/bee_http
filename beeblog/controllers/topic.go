package controllers

import (
	"bee_http/beeblog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"

	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}
}

func (c *TopicController) Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	title := c.Input().Get("title")
	content := c.Input().Get("content")

	var err error
	err = models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic", 302)
}

func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}

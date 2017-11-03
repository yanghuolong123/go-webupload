package controllers

import (
	"yhl/help"
)

type HomeController struct {
	help.BaseController
}

func (this *HomeController) Get() {
	help.Log.Info(this.Ctx.Input.Domain())
	this.SendRes(0, "success", nil)
}

func (this *HomeController) Post() {
	this.SendRes(0, "success", "post")
}

package controllers

import (
	"runtime"
	"yhl/help"
)

type HomeController struct {
	help.BaseController
}

func (this *HomeController) Get() {
	help.Log.Info(this.Ctx.Input.Domain())
	cpunum := runtime.NumCPU()
	this.SendRes(0, "success", cpunum)
}

func (this *HomeController) Post() {
	this.SendRes(0, "success", "post")
}

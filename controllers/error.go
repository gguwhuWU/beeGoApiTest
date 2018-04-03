package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.Data["json"] = map[string]string{"msg": "page not found", "status": "error404"}
	this.ServeJSON()
}

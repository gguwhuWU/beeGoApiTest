package controllers

import (
	"beeApiTest/models"

	"github.com/astaxie/beego"
)

type ClassController struct {
	beego.Controller
}

func (this *ClassController) GetById() {
	classId := this.Ctx.Input.Param(":classId")
	if classId != "" {
		class, err := models.GetClassById(classId)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = class
		}
	}
	this.ServeJSON()
}

func (this *ClassController) GetByName() {
	className := this.Ctx.Input.Param(":className")
	if className != "" {
		class, err := models.GetClassByName(className)
		if err != nil {
			this.Data["json"] = map[string]string{"status": err.Error()}
		} else {
			this.Data["json"] = class
		}
	}
	this.ServeJSON()
}

func (this *ClassController) GetAllClass() {
	classes := models.GetAllClasses()
	this.Data["json"] = classes
	this.ServeJSON()
}

func (this *ClassController) PostNewClass() {
	var newClass models.Class
	this.ParseForm(&newClass)
	newClassId := models.AddNewClass(&newClass)
	this.Data["json"] = map[string]string{"classId": newClassId}
	this.ServeJSON()
}

package controllers

import (
	"beeApiTest/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type StudentController struct {
	beego.Controller
}

func (s *StudentController) Get() {

	StudentId := s.Ctx.Input.Param(":studentId")
	if StudentId != "" {
		st, err := models.GetStudent(StudentId)
		if err != nil {
			s.Data["json"] = err.Error()
		} else {
			s.Data["json"] = st
		}
	}
	s.ServeJSON()
}

func (s *StudentController) GetAll() {
	// s.Redirect("/", 302)
	s.Abort("404")

	sts := models.GetAllStudents()
	s.Data["json"] = sts
	s.ServeJSON()
}

func (this *StudentController) PostNewStudent() {
	var newStudent models.Student
	this.ParseForm(&newStudent)
	newStudentId := models.AddNewStudent(&newStudent)
	this.Data["json"] = map[string]string{"StudentId": newStudentId}
	this.ServeJSON()
}

func (this *StudentController) DelStudentByTID() {
	// sid := this.Ctx.Input.Param(":studentId")
	sid := this.GetString(":studentId")
	models.DeleteStudent(sid)
	this.Data["json"] = map[string]string{"status": "success", "data": sid}
	this.ServeJSON()
}

type Demo struct {
	Status  string `json:stauts`
	IsAdult bool   `json:isAdult`
}

func (this *StudentController) PutRequestBodyTest() {
	newD := Demo{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &newD)
	this.Ctx.ResponseWriter.Header().Set("ServerMan", "StockinTea")
	this.Data["json"] = newD
	this.ServeJSON()
}

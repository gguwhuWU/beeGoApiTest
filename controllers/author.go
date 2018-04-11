package controllers

import (
	"beeApiTest/models"
	"beeApiTest/repository"
	"strconv"

	"github.com/astaxie/beego"
)

type AuthorController struct {
	beego.Controller
}

func (controller *AuthorController) Get() {

	AuthorId := controller.Ctx.Input.Param(":authorId")
	if AuthorId != "" {
		number, _ := strconv.Atoi(AuthorId)
		r, err := repository.GetAuthor(number)

		if err != nil {
			controller.Data["json"] = err.Error()
		} else {
			controller.Data["json"] = r
		}
	}

	controller.ServeJSON()
}

func (controller *AuthorController) GetAll() {
	r := repository.GetAllAuthors()
	controller.Data["json"] = r
	controller.ServeJSON()
}

func (controller *AuthorController) PostNewAuthor() {
	var newAuthor models.Author
	controller.ParseForm(&newAuthor)
	newAuthorId := repository.AddAuthor(&newAuthor)
	controller.Data["json"] = map[string]int64{"authorId": newAuthorId}
	controller.ServeJSON()
}

func (controller *AuthorController) DelAuthorById() {
	aid := controller.GetString(":authorId")
	number, _ := strconv.Atoi(aid)
	repository.DeleteAuthor(number)
	controller.Data["json"] = map[string]string{"status": "success", "data": aid}
	controller.ServeJSON()
}

func (controller *AuthorController) UpAuthorById() {
	var newAuthor models.Author
	aid := controller.GetString(":authorId")
	number, _ := strconv.Atoi(aid)
	controller.ParseForm(&newAuthor)
	r, err := repository.UpdateAuthor(number, &newAuthor)

	if err != nil {
		controller.Data["json"] = err.Error()
	} else {
		controller.Data["json"] = r
	}
	controller.ServeJSON()
}

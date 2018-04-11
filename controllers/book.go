package controllers

import (
	"github.com/astaxie/beego"
)

type BookController struct {
	beego.Controller
}

// func (b *BookController) Get() {

// 	BookId := b.Ctx.Input.Param(":bookId")
// 	if BookId != "" {
// 		number, _ := strconv.Atoi(BookId)
// 		r, err := models.GetBook(number)

// 		if err != nil {
// 			b.Data["json"] = err.Error()
// 		} else {
// 			b.Data["json"] = r
// 		}
// 	}

// 	b.ServeJSON()
// }

// func (b *BookController) GetAll() {
// 	r := models.GetAllBooks()
// 	b.Data["json"] = r
// 	b.ServeJSON()
// }

// func (b *BookController) PostNewBook() {
// 	var newBook models.Book
// 	b.ParseForm(&newBook)
// 	newBookId := models.AddBook(&newBook)
// 	b.Data["json"] = map[string]int64{"bookId": newBookId}
// 	b.ServeJSON()
// }

// func (b *BookController) DelBookById() {
// 	bid := b.GetString(":bookId")
// 	number, _ := strconv.Atoi(bid)
// 	models.DeleteBook(number)
// 	b.Data["json"] = map[string]string{"status": "success", "data": bid}
// 	b.ServeJSON()
// }

// func (b *BookController) UpBookById() {

// }

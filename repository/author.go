package repository

import (
	"beeApiTest/models"
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "wu:wu@/testwu?charset=utf8", 30)
	orm.RegisterModel(new(models.Book), new(models.Author), new(models.Publisher), new(models.Taxonomy))

	db := orm.NewOrm()
	//Database alias
	name := "default"
	// Drop table and re-create.
	force := false
	// Print log.
	verbose := true
	err := orm.RunSyncdb(name, force, verbose) //this is to create/drops the tables
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in SyncDB")
	}
	db.Using("testwu")
}

func AddAuthor(a *models.Author) int64 {
	o := orm.NewOrm()

	// insert
	//author := models.Author{Name: "wu", Email: "wu1231dfgdfg123@gmail.com"}
	id, err := o.Insert(a)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	return id
}

func GetAuthor(id int) (r *models.Author, err error) {
	o := orm.NewOrm()
	author := models.Author{ID: id}

	err = o.Read(&author)

	return &author, err
}

func GetAllAuthors() []*models.Author {
	var authors []*models.Author
	o := orm.NewOrm()
	authorModel := new(models.Author)
	o.QueryTable(authorModel).All(&authors)

	return authors
}

func UpdateAuthor(id int, a *models.Author) (r *models.Author, err error) {
	o := orm.NewOrm()
	author := models.Author{ID: id}

	if o.Read(&author) == nil {
		author.Name = a.Name
		author.Email = a.Email

		if num, err := o.Update(&author); err == nil {
			fmt.Println(num)
		}

		return &author, err
	}

	return nil, errors.New("User Not Exist")
}

func DeleteAuthor(id int) {
	o := orm.NewOrm()
	if num, err := o.Delete(&models.Author{ID: id}); err == nil {
		fmt.Println(num)
	}
}

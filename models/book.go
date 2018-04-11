package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

type Book struct {
	ID          int `orm:"auto"`
	Name        string
	ISBN        string      `orm:"unique;size(100)"`
	PublishTime time.Time   `orm:"type(date)"`
	Price       int         `orm:"null"`
	Page        int         `orm:"default(1)"`
	Lang        string      `orm:"null"`
	CreatedTime time.Time   `orm:"auto_now_add;type(datetime)"`
	Publisher   *Publisher  `orm:"rel(fk)"`  // RelForeignKey relation
	Authors     []*Author   `orm:"rel(m2m)"` // ManyToMany relation
	Taxonomies  []*Taxonomy `orm:"rel(m2m)"` // ManyToMany relation
}

type Author struct {
	ID          int `orm:"auto"`
	Name        string
	Email       string
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	Books       []*Book   `orm:"reverse(many)"`
}

type Publisher struct {
	ID          int `orm:"auto"`
	Name        string
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	Books       []*Book   `orm:"reverse(many)"` // reverse relationship of fk
}

type Taxonomy struct {
	ID          int `orm:"auto"`
	Name        string
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	Books       []*Book   `orm:"reverse(many)"` // reverse relationship of fk
}

// multiple fields index
func (b *Book) TableIndex() [][]string {
	return [][]string{
		[]string{"ISBN"},
		[]string{"PublishTime"},
		[]string{"Name"},
		[]string{"Price"},
	}
}

// multiple fields index
func (a *Author) TableIndex() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

// multiple fields index
func (p *Publisher) TableIndex() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

// multiple fields index
func (t *Taxonomy) TableIndex() [][]string {
	return [][]string{
		[]string{"Name"},
	}
}

func (b *Book) TableName() string {
	return "Books"
}

func (a *Author) TableName() string {
	return "Authors"
}

func (p *Publisher) TableName() string {
	return "Publishers"
}

func (t *Taxonomy) TableName() string {
	return "Taxonomies"
}

// func init() {
// 	orm.RegisterDriver("mysql", orm.DRMySQL)
// 	orm.RegisterDataBase("default", "mysql", "wu:wu@/testwu?charset=utf8", 30)
// 	orm.RegisterModel(new(Book), new(Author), new(Publisher), new(Taxonomy))

// 	db := orm.NewOrm()
// 	//Database alias
// 	name := "default"
// 	// Drop table and re-create.
// 	force := false
// 	// Print log.
// 	verbose := true
// 	err := orm.RunSyncdb(name, force, verbose) //this is to create/drops the tables
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("Error in SyncDB")
// 	}
// 	db.Using("testwu")
// }

// func AddBook(b *Book) int64 {
// 	o := orm.NewOrm()

// 	// insert
// 	// publisher := Publisher{Name: "jojo~3"}
// 	// id, err := o.Insert(&publisher)
// 	// fmt.Printf("ID: %d, ERR: %v\n", id, err)

// 	// author := Author{Name: "wu", Email: "wu1231dfgdfg123@gmail.com"}
// 	// id2, err2 := o.Insert(&author)
// 	// fmt.Printf("ID: %d, ERR: %v\n", id2, err2)

// 	// taxonomy := Taxonomy{Name: "animal"}
// 	// id3, err3 := o.Insert(&taxonomy)
// 	// fmt.Printf("ID: %d, ERR: %v\n", id3, err3)

// 	publisher := Publisher{ID: 1}
// 	o.Read(&publisher)

// 	author := Author{ID: 1}
// 	o.Read(&author)
// 	var authors []*Author
// 	authors[0] = &author

// 	taxonomy := Taxonomy{ID: 1}
// 	o.Read(&taxonomy)
// 	var taxonomies []*Taxonomy
// 	taxonomies[0] = &taxonomy

// 	book := Book{Name: "slene", ISBN: "9864342924", PublishTime: time.Date(
// 		2009, 11, 17, 20, 34, 58, 651387237, time.UTC), Price: 350, Page: 220, Publisher: &publisher, Authors: authors, Taxonomies: taxonomies}

// 	// insert
// 	id, err := o.Insert(&book)
// 	fmt.Printf("ID: %d, ERR: %v\n", id, err)

// 	return id
// }

// func GetBook(id int) (r *Book, err error) {

// 	return nil, errors.New("User not exists")
// }

// func GetAllBooks() map[string]*Book {
// 	return make(map[string]*Book)
// }

// func UpdateBook(id int, b *Book) (r *Book, err error) {

// 	return nil, errors.New("User Not Exist")
// }

// func DeleteBook(id int) {
// 	//delete(UserList, uid)
// }

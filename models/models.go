package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// create orm object, and it will use `default` database

//	o := orm.NewOrm()
// data
//	user := new(User)
//	user.UserName = "mike"

// insert data
//	id, err := o.Insert(user)
//	if err != nil {
//		logs.Info(err)
//	}
//	_ = id

// Book -
type Book struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
	Type int    `orm:"column(type)"`
	Isbn string `orm:"column(isbn)"`
	Desc string `orm:"column(desc)"`
	Lang string `orm:"column(lang)"`
	Item string `orm:"column(item)"`
}

// BookShelf -
type BookShelf struct {
	ID        int `orm:"column(id)"`
	UserId    int `orm:"column(user_id)"`
	BookId    int `orm:"column(book_id)"`
	BookState int `orm:"column(isbn)"`
	Progress  int `orm:"column(desc)"`
}

func (bs *BookShelf) TableName() string {
	return "book_shelf"
}

// Circle -
type Circle struct {
	ID        int    `orm:"column(id)"`
	CreatorId int    `orm:"column(creator_id)"`
	title     string `orm:"column(title)"`
}

// Discuss -
type Discuss struct {
	ID       int    `orm:"column(id)"`
	CircleId int    `orm:"column(circle_id)"`
	time     string `orm:"column(time)"`
	FromId   int    `orm:"column(from_id)"`
	ToId     int    `orm:"column(to_id)"`
	Content  string `orm:"column(content)"`
}

// Punch -
type Punch struct {
	ID        int    `orm:"column(id)"`
	UserId    int    `orm:"column(user_id)"`
	BookId    int    `orm:"column(book_id)"`
	TimePoint string `orm:"column(timepoint)"`
	ReadTime  int    `orm:"column(treadtime)"`
}

// Review -
type Review struct {
	ID       int    `orm:"column(id)"`
	UserId   int    `orm:"column(user_id)"`
	CircleId int    `orm:"column(circle_id)"`
	Content  string `orm:"column(content)"`
	time     string `orm:"column(time)"`
}

// Experience -
type Experience struct {
	ID      int    `orm:"column(id)"`
	UserId  int    `orm:"column(user_id)"`
	BookId  int    `orm:"column(book_id)"`
	Content string `orm:"column(content)"`
}

// User -
type User struct {
	ID        int    `orm:"column(id)" json:"id"`
	Number    string `orm:"column(number)" json:"number"`
	UserName  string `orm:"column(username)" json:"username"`
	Password  string `orm:"column(password)" json:"password"`
	Type      string `orm:"column(type)" json:"type"`
	CircleIds string `orm:"column(circle_ids)" json:"circle_ids"`
}

// type User struct {
// 	ID        int    `orm:"column(id)"`
// 	Number    string `orm:"column(number)"`
// 	UserName  string `orm:"column(username)"`
// 	Password  string `orm:"column(password)"`
// 	Type      string `orm:"column(type)"`
// 	CircleIds string `orm:"column(circle_ids)"`
// }

func init() {
	// need to register models in init
	orm.RegisterModel(new(Book))
	orm.RegisterModel(new(BookShelf))
	orm.RegisterModel(new(Circle))
	orm.RegisterModel(new(Discuss))
	orm.RegisterModel(new(Punch))
	orm.RegisterModel(new(Review))
	orm.RegisterModel(new(Experience))
	orm.RegisterModel(new(User))
	// need to register db driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/libsys?charset=utf8")

	// automatically build table
	orm.RunSyncdb("default", false, true)
}

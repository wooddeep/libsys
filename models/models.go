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

type Base struct {
}

// Book -
type Book struct {
	ID   int    `orm:"column(id)" json:"id"`
	Name string `orm:"column(name)" json:"name"`
	Type int    `orm:"column(type)" json:"type"`
	Isbn string `orm:"column(isbn)" json:"isbn"`
	Desc string `orm:"column(desc)" json:"desc"`
	Lang string `orm:"column(lang)" json:"lang"`
	Item string `orm:"column(item)" json:"item"`
}

// BookShelf -
type BookShelf struct {
	Base
	ID        int `orm:"column(id)" json:"id"`
	UserId    int `orm:"column(user_id)" json:"user_id"`
	BookId    int `orm:"column(book_id)" json:"book_id"`
	BookState int `orm:"column(book_state)" json:"book_state"`
	Progress  int `orm:"column(progress)" json:"progress"`
}

func (bs *BookShelf) TableName() string {
	return "book_shelf"
}

// Circle -
type Circle struct {
	ID        int    `orm:"column(id)" json:"id"`
	CreatorId int    `orm:"column(creator_id)" json:"creator_id"`
	title     string `orm:"column(title)" json:"title"`
}

// Discuss -
type Discuss struct {
	ID       int    `orm:"column(id)" json:"id"`
	CircleId int    `orm:"column(circle_id)" json:"circle_id"`
	time     string `orm:"column(time)" json:"time"`
	FromId   int    `orm:"column(from_id)" json:"from_id"`
	ToId     int    `orm:"column(to_id)" json:"to_id"`
	Content  string `orm:"column(content)" json:"content"`
}

// Punch -
type Punch struct {
	ID        int    `orm:"column(id)" json:"id"`
	UserId    int    `orm:"column(user_id)" json:"user_id"`
	BookId    int    `orm:"column(book_id)" json:"book_id"`
	TimePoint string `orm:"column(timepoint)" json:"timepoint"`
	ReadTime  int    `orm:"column(treadtime)" json:"treadtime"`
}

// Review -
type Review struct {
	ID       int    `orm:"column(id)" json:"id"`
	UserId   int    `orm:"column(user_id)" json:"user_id"`
	CircleId int    `orm:"column(circle_id)" json:"circle_id"`
	Content  string `orm:"column(content)" json:"content"`
	time     string `orm:"column(time)" json:"time"`
}

// Experience -
type Experience struct {
	ID      int    `orm:"column(id)" json:"id"`
	UserId  int    `orm:"column(user_id)" json:"user_id"`
	BookId  int    `orm:"column(book_id)" json:"book_id"`
	Content string `orm:"column(content)" json:"content"`
}

// User -
type User struct {
	ID        int    `orm:"column(id)" json:"id"`
	Number    string `orm:"column(number)" json:"number"`
	UserName  string `orm:"column(username)" json:"username"`
	Password  string `orm:"column(password)" json:"password"`
	Type      string `orm:"column(type)" json:"type"`
	CircleIds string `orm:"column(circle_ids)" json:"circle_ids"`
	Name      string `orm:"column(name)" json:"name"`
}

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

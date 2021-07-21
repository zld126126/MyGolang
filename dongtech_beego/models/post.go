package models

type Post struct {
	Id    int    `orm:"auto"`
	Title string `orm:"size(100)"`
	User  *User  `orm:"rel(fk)"`
}

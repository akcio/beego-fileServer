package models

import (
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type User struct {
	Id int64 `orm:"auto"`
	Login string `orm:"size(100)"`
	Password string `orm:"size(100)"`
}

type File struct {
	Id int64 `orm:"auto"`
	Hash1 string
	Hash2 string
	Stored string
}

type UserFile struct {
	LinkId int64 `orm:"auto"`
	UserId int64
	FileId int64
	UserFileName string
	Mode int8
	UploadTime string
}

func (u *User) TableName() string {
	return "users"
}

func (f * File) TableName() string {
	return "files"
}

func (uf * UserFile) TableName() string {
	return "users_files"
}

func init()  {
	orm.Debug = true
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/test.db")

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(File))
	orm.RegisterModel(new(UserFile))
}

func GetGUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func GetSalt() string {
	return "#1234!_22query=+"
}
package controllers

import (
	"crypto"
	"crypto/md5"
	"encoding/hex"
	"beego-fileServer/models"
	"beego-fileServer/models/helpers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"
)

type FilesController struct {
	beego.Controller
}


func (this *FilesController) Index() {
	this.Data["Error"] = "user"
	this.TplName = "forms/login.tpl"
}

func getFileExtension(fileName string) string {
	res := ""
	canWrite := false
	for i:=0; i < len(fileName); i++ {
		if fileName[i] == '.' {
			canWrite = true
		}
		if canWrite {
			res += string(fileName[i])
		}
	}
	return res
}

type Test struct {
	UserId int64
	FileId int64
	UserFileName string
	Mode int
	UploadTime string
	Stored string
}

func (this *FilesController) List() {
	helpers.SetLayoutFor(&this.Controller)
	var user = helpers.GetCurrentUser(&this.Controller)
	var o = helpers.GetORM()
	this.TplName = "fileList.tpl"
	//this.Data["items"] = user.Files

	//var files []*models.File
	//o.QueryTable(new(models.UserFile)).Filter("user_id", user.Id).OrderBy("upload_time").All(&files)
	//
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(
		"users_files.user_id",
		"users_files.file_id",
		"users_files.user_file_name",
		"users_files.mode",
		"users_files.upload_time",
		"files.stored").
		From("users_files").
		InnerJoin("files").On("users_files.file_id = files.id").
		Where("users_files.user_id = ?").
		OrderBy("users_files.upload_time")

	var userFiles []Test

	sql := qb.String()
	o.Raw(sql, user.Id).QueryRows(&userFiles)

	this.Data["Files"] = userFiles
	this.Data["Val"] = len(userFiles)
}

func (this *FilesController)  Download() {
	helpers.SetLayoutFor(&this.Controller)

	fileMarker := this.Ctx.Input.Param(":name")
	this.TplName = "index.tpl"
	this.Data["Website"] = fileMarker

	var o = helpers.GetORM()
	var user = helpers.GetCurrentUser(&this.Controller)
	//this.Data["items"] = user.Files

	//var files []*models.File
	//o.QueryTable(new(models.UserFile)).Filter("user_id", user.Id).OrderBy("upload_time").All(&files)
	//
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(
		"users_files.user_id",
		"users_files.file_id",
		"users_files.user_file_name",
		"users_files.mode",
		"users_files.upload_time",
		"files.stored").
		From("users_files").
		InnerJoin("files").On("users_files.file_id = files.id").
		Where("users_files.user_id = ? AND users_files.upload_time = ?")

	var userFiles []Test
	sql := qb.String()
	o.Raw(sql, user.Id, fileMarker).QueryRows(&userFiles)
	this.Data["Website"] = len(userFiles)
	if len(userFiles) > 0 {
		this.Data["Email"] = userFiles[0].Stored
		this.Ctx.Output.Download(userFiles[0].Stored)
	}

	this.Redirect("/", 302)
}

func (this *FilesController) DeleteLink() {

	fileMarker := this.Ctx.Input.Param(":name")
	this.TplName = "index.tpl"
	this.Data["Website"] = fileMarker
	var o = helpers.GetORM()
	var user = helpers.GetCurrentUser(&this.Controller)
	if err := o.QueryTable(new(models.User)).Filter("id", user.Id).One(&user); err != nil {
		this.Redirect("/login", 302)
	}
	//this.Data["items"] = user.Files

	//var files []*models.File
	//o.QueryTable(new(models.UserFile)).Filter("user_id", user.Id).OrderBy("upload_time").All(&files)
	//
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(
	"users_files.user_id",
	"users_files.file_id",
	"users_files.user_file_name",
	"users_files.mode",
	"users_files.upload_time",
	"files.stored").
	From("users_files").
	InnerJoin("files").On("users_files.file_id = files.id").
	Where("users_files.user_id = ? AND users_files.upload_time = ?")

	var userFiles []Test
	sql := qb.String()
	o.Raw(sql, user.Id, fileMarker).QueryRows(&userFiles)
	this.Data["Website"] = len(userFiles)
	if len(userFiles) > 0 {
		_, err := o.Raw("DELETE FROM users_files " +
			"WHERE user_id = ? AND file_id = ? AND upload_time = ?",
			userFiles[0].UserId,
			userFiles[0].FileId,
			userFiles[0].UploadTime).Exec()
		this.Data["Error"] = err
	}
	this.Redirect("/user/"+user.Login, 302)
}

func (this *FilesController) Upload() {
	helpers.SetLayoutFor(&this.Controller)
	//var o = helpers.GetORM()
	var user = helpers.GetCurrentUser(&this.Controller)
	this.TplName = "forms/upload.tpl"

	if this.Ctx.Request.Method == "GET" {
		return
	}
	var files, err = this.GetFiles("the_file")

	if err != nil {
		return
	}
	for _, header:= range files {
		ProcessFile(&user, header)
	}

	this.Redirect("/user/"+user.Login, 302)
}

func ProcessFile(user *models.User, header *multipart.FileHeader) {
	var fileFolder = beego.AppConfig.DefaultString("filesFolder", "/tmp/files/")

	file, errOpen := header.Open()
	if errOpen != nil {
		return
	}
	newFileData, _ := ioutil.ReadAll(file)
	var o = helpers.GetORM()

	var hasherMd5 = md5.New()
	var hasherSha256 = crypto.SHA256.New()
	hasherMd5.Write(newFileData)
	hasherSha256.Write(newFileData)

	md5Hash := hex.EncodeToString(hasherMd5.Sum(nil))
	sha256Hash := hex.EncodeToString(hasherSha256.Sum(nil))


	var fileItem models.File
	err := o.QueryTable(new(models.File)).Filter("hash1", md5Hash).Filter("hash2", sha256Hash).One(&fileItem)
	if err == orm.ErrNoRows{
		newFileNamePath := fileFolder + models.GetGUID() +  getFileExtension(header.Filename)
		f, err2 := os.OpenFile(newFileNamePath, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0777)
		if err2 != nil {
			return
		}
		defer f.Close()
		io.Copy(f, file)
		fileToAdd := models.File{Hash1:md5Hash, Hash2:sha256Hash, Stored:newFileNamePath}
		//this.Data["Message"] = "No file in database"
		if id, err := o.Insert(&fileToAdd); err == nil{
			fileToAdd.Id = id
			//this.Data["Message"] = "file created"
			AddFileToUser(user, &fileToAdd, header.Filename)
			//this.Redirect("/user/" + user.Login, 302)
			return
		}
	}
	if err == orm.ErrMultiRows {
		//this.Data["Error"] = "server error"
		//os.Remove(newFileNamePath)
		return
	}

	AddFileToUser(user, &fileItem, header.Filename)
}

func AddFileToUser(user *models.User, file *models.File, fileName string) {
	var o = helpers.GetORM()

	var link = models.UserFile{UserId:user.Id, FileId:file.Id, UserFileName:fileName, Mode:0, UploadTime: time.Now().Format(time.RFC3339)}
	o.Insert(&link)
}

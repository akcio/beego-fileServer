package controllers

import (
	"crypto"
	"crypto/md5"
	"encoding/hex"
	"fileServer/testBeeGo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
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
		if canWrite {
			res += string(fileName[i])
		}
		if fileName[i] == '.' {
			canWrite = true
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
	sess := this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	userId := sess.Get("userId")
	o := orm.NewOrm()
	o.Using("default")
	var user models.User

	if err := o.QueryTable(new(models.User)).Filter("id", userId).One(&user); err != nil {
		this.Redirect("/login", 302)
	}
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
	o.Raw(sql, userId).QueryRows(&userFiles)

	this.Data["Files"] = userFiles
	this.Data["Val"] = len(userFiles)
}


func (this *FilesController) Upload() {

	this.TplName = "forms/upload.tpl"

	var sess = this.StartSession()
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	var userId = sess.Get("userId")
	var o = orm.NewOrm()
	o.Using("default")
	var user models.User

	if err := o.QueryTable(new(models.User)).Filter("id", userId).One(&user); err != nil {
		this.Redirect("/login", 302)
	}

	var file, header, _ = this.GetFile("the_file")

	if file == nil {
		return
	}

	//var p []byte

	newFileNamePath := "/tmp/files/"+ models.GetGUID() + "." +  getFileExtension(header.Filename)
	file.Close()
	this.SaveToFile("the_file", newFileNamePath)

	var newFile, _ = ioutil.ReadFile(newFileNamePath)

	var hasherMd5 = md5.New()
	var hasherSha256 = crypto.SHA256.New()
	hasherMd5.Write(newFile)
	hasherSha256.Write(newFile)

	md5Hash := hex.EncodeToString(hasherMd5.Sum(nil))
	sha256Hash := hex.EncodeToString(hasherSha256.Sum(nil))


	//file.Close()
	this.Data["Error"] = newFileNamePath
	this.Data["Hash1"] = md5Hash
	this.Data["Hash2"] = sha256Hash
	//
	var fileItem models.File
	err := o.QueryTable(new(models.File)).Filter("hash1", md5Hash).Filter("hash2", sha256Hash).One(&fileItem)
	if err == orm.ErrNoRows{
		fileToAdd := models.File{Hash1:md5Hash, Hash2:sha256Hash, Stored:newFileNamePath}
		this.Data["Message"] = "No file in database"
		if id, err := o.Insert(&fileToAdd); err == nil{
			fileToAdd.Id = id
			this.Data["Message"] = "file created"
			AddFileToUser(&user, &fileToAdd, header.Filename)
			this.Redirect("/upload", 302)
			return
		}
	}
	if err == orm.ErrMultiRows {
		this.Data["Error"] = "server error"
		os.Remove(newFileNamePath)
		return
	}

	AddFileToUser(&user, &fileItem, header.Filename)
	os.Remove(newFileNamePath)
}

func AddFileToUser(user *models.User, file *models.File, fileName string) {
	var o = orm.NewOrm()
	o.Using("default")

	var link = models.UserFile{UserId:user.Id, FileId:file.Id, UserFileName:fileName, Mode:0, UploadTime: time.Now().String()}
	o.Insert(&link)
}

package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"os"
	"shujvrenzhengzhichonglai/models"
	"strings"
	"time"
)

type UploadFileController struct {
	beego.Controller
}

func (u *UploadFileController)Post(){
	title:=u.Ctx.Request.PostFormValue("upload_title")
	fmt.Println("电子数据标签",title)
	file,header,err:=u.GetFile("yuhongwei")
	if err != nil {
		u.Ctx.WriteString("抱歉，文件解析失败，请重试")
		return
	}
	defer file.Close()

	saveFilePath:="static/upload"+header.Filename
	saveFile,err:=os.OpenFile(saveFilePath,os.O_CREATE|os.O_RDWR,777)
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重试！")
		return
	}
	_,err=io.Copy(saveFile,file)
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重新尝试！")
		return
	}
	hash256:=sha256.New()
	fileBytes,_:=ioutil.ReadAll(file)
	hash256.Write(fileBytes)
	hashBytes:=hash256.Sum(nil)
	fmt.Println(hex.EncodeToString(hashBytes))

	//u.Ctx.WriteString("恭喜，已接收到上传文件")
	user1, err := models.User{Phone: phone}.QueryUserByPhone()
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请稍后再试!")
		return
	}

	//把上传的文件作为记录保存到数据库当中
	//① 计算md5值
	md5Hash := md5.New()
	fileMd5Bytes, err := ioutil.ReadAll(saveFile)
	md5Hash.Write(fileMd5Bytes)
	bytes := md5Hash.Sum(nil)
	record := models.UploadRecord{
		UserId:    user1.Id,
		FileName:  header.Filename,
		FileSize:  header.Size,
		FileCert:  hex.EncodeToString(bytes),
		FileTitle: tile,
		CertTime:  time.Now().Unix(),
	}
	//② 保存认证数据到数据库中
	_, err = record.SaveRecord()
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证保存失败，请稍后再试!")
		return
	}
	//上传文件保存到数据库中成功
	records, err := models.QueryRecordsByUserId(user1.Id)
	if err != nil {
		u.Ctx.WriteString("抱歉, 获取电子数据列表失败, 请重新尝试!")
		return
	}
	u.Data["Records"] = records
	u.TplName = "list_record.html"
}

func (u *UploadFileController)Post1(){
	title:=u.Ctx.Request.PostFormValue("upload_title")

	file,header,err:=u.GetFile("yuhongwei")
	if err != nil {
		u.Ctx.WriteString("抱歉，文件解析失败，请重试")
		return
	}
	defer file.Close()

	fmt.Println("自定义的标签",title)

	fmt.Println("上传文件的名称",header.Filename)

	fileNameSlice:=strings.Split(header.Filename,".")
	fileType:=fileNameSlice[1]
	fmt.Println(fileNameSlice)
	fmt.Println(":",strings.TrimSpace(fileType))
	isJpg:=strings.HasSuffix(header.Filename,".jpg")
	isPng:=strings.HasSuffix(header.Filename,".png")
	if !isJpg&&!isPng{
		u.Ctx.WriteString("抱歉，文件类型不符合，请上传符合格式的文件")
		return
	}
	config:=beego.AppConfig
	fileSize,err:=config.Int64("file_size")

	if header.Size/1024>fileSize{
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
		return
	}
	fmt.Println("上传文件的大小：",header.Size)

	saveDir:="static/upload"
	_,err=os.Open(saveDir)

	if err!=nil{
		err=os.Mkdir(saveDir,777)
		if err != nil {
			fmt.Println(err.Error())
			u.Ctx.WriteString("抱歉，文件认证遇到错误，请重试")
			return
		}
	}
	saveName:=saveDir+"/"+header.Filename
	fmt.Println("要保存的文件名",saveName)
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，文件认证失败，请重试")
		return
	}
	fmt.Println("上传的文件",file )
	u.Ctx.WriteString("已获取上传文件")
}

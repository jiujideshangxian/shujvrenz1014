package models

import "shujvrenzhengzhichonglai/db_mysql"

type UploadRecord struct {
	Id        int
	UserId    int
	FileName  string
	FileSize  int64
	FileCert  string
	FileTitle string
	CertTime  int64
}

func (u UploadRecord)SaveRecord()(int64,error){
    rs,err:=db_mysql.Db.Exec("insert into upload_record(user_id,file_name,file_size,file_cert,file_title,cert_time)"+
    	"values(?,?,?,?,?,?)")
    if err!=nil{
    	return -1, err
	}
	id,err:=rs.RowsAffected()
	if err!=nil{
		return -1,err
	}
	return id,nil
}

func QueryRecordsByUserId(userId int) ([]UploadRecord,error) {
	rs,err:=db_mysql.Db.Query("select id,user_id,file_name,file_size,file_cert,file_title,cert_time) where user_id=?",userId)
	if err != nil {
		return nil, err
	}
	records:=make([]UploadRecord,0)
	for rs.Next() {
		var record UploadRecord
		err:=rs.Scan(&record.Id,&record.UserId,&record.FileName,&record.FileSize,&record.FileCert, &record.FileTitle, &record.CertTime)
		if err != nil {
			return nil,err
		}
		records=append(records,record)
	}
	return records,nil
}

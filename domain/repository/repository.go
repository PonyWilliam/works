package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"work/domain/model"
	"work/common"
)

type IWorker interface {
	InitTable() error
	CreateWorker(worker *model.Workers) (int64,error)
	UpdateWorker(worker *model.Workers) (int64,error)
	DeleteWorkerByID(int64) error
	FindWorkerByID(int64)(model.Workers,error)
	FindWorkersByName(string)([]model.Workers,error)
	FindAll()([]model.Workers,error)
	Login(username string,password string)(bool,string,error)
}
func NewWorkerRepository(db *gorm.DB)IWorker{
	return &WorkersRepository{mysqlDB: db}
}
type WorkersRepository struct{
	mysqlDB *gorm.DB
}
func (w *WorkersRepository) InitTable() error{
	if w.mysqlDB.HasTable(&model.Workers{}){
		return nil
	}
	return w.mysqlDB.CreateTable(&model.Workers{}).Error
}
func (w *WorkersRepository) CreateWorker(worker *model.Workers) (int64,error){
	return worker.ID,w.mysqlDB.Model(worker).Create(&worker).Error
}
func (w *WorkersRepository) UpdateWorker(worker *model.Workers) (int64,error){
	return worker.ID,w.mysqlDB.Model(worker).Update(&worker).Error
}
func (w *WorkersRepository) DeleteWorkerByID(id int64) error{
	return w.mysqlDB.Where("id = ?",id).Delete(&model.Workers{}).Error
}
func (w *WorkersRepository) FindWorkerByID(id int64) (worker model.Workers,err error){
	return worker,w.mysqlDB.Model(&model.Workers{}).Where("id  = ?",id).Find(&worker).Error
}
func (w *WorkersRepository) FindWorkersByName(name string) (workers []model.Workers,err error){
	return workers,w.mysqlDB.Model(&model.Workers{}).Where("name  = ?",name).Find(&workers).Error
}
func (w *WorkersRepository) FindAll() (workers []model.Workers,err error){
	return workers,w.mysqlDB.Model(&model.Workers{}).Find(&workers).Error
}
func (w *WorkersRepository) Login(username string,password string)(bool,string,error){
	works := &model.Workers{}
	err := w.mysqlDB.Model(&model.Workers{}).Find(&works).Where("username = ?",username).Error
	if err!=nil{
		return false,"",err
	}
	if password == works.HashPassword{
		return false,"",fmt.Errorf("密码错误")
	}
	//还要创建token方便下次登陆,通过service发给客户浏览器
	token,err := common.CreateToken(username,password)
	if err != nil{
		return false,"",fmt.Errorf(err.Error())
	}
	return true,token,nil
}
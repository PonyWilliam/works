/*
		srv := service.New(
		service.Name("work"),
		service.Version("latest"),
		)
	_ = srv.Handle(new(handler.Work))
	if err:=srv.Run();err!=nil{
		logger.Error(err)
		logger.Error("1234")
	}
*/package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"strconv"
	"work/domain/repository"
	services2 "work/domain/services"
	"work/handler"
	pb "work/proto"
)
type Info struct{
	Host string
	Port int64
	User string
	Pwd string
	DataBase string
}
func main() {
	var mysqlInfo = Info{
		Host: "nc.lllui.cn",
		Port: 3306,
		User: "gostudy",
		Pwd: "gostudy",
		DataBase: "gostudy",
	}
	//consulConfig,err := common.GetConsualConfig("127.0.0.1",8500,"/micro/config")
	//配置中心
	/*if err != nil{
		logger.Fatal(err)
	}*/

	//注册中心
	/*consulRegistry := consul.NewRegistry(
		func(options *registry.Options){
			options.Addrs = []string{"127.0.0.1"}
			options.Timeout = time.Second * 10
		},
	)*/
	srv := service.New(
		service.Name("work"),
		service.Version("latest"),
		//service.HandleSignal(true),
		//service.Address("127.0.0.1:8083"),

	)
	//_ = srv.Handle(services2.InitRouters())
	db,err := gorm.Open("mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host + ":"+ strconv.FormatInt(mysqlInfo.Port,10) +")/"+mysqlInfo.DataBase+"?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil{
		logger.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)
	srv.Init()
	rp := repository.NewWorkerRepository(db)
	err =rp.InitTable()
	if err!=nil{
		err := rp.InitTable()
		if err!=nil{
			logger.Error(err)
		}
	}

	WorkServices := services2.NewWorkerServices(repository.NewWorkerRepository(db))
	//err = works.RegisterWorksHandler(srv.Server(),&handler.Works{WorkService:WorkServices})
	_ = pb.RegisterWorkHandler(srv.Server(), &handler.Work{WorkService: WorkServices})
	if err:=srv.Run();err!=nil{
		logger.Fatal(err)
	}
}
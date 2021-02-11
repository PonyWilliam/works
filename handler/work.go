package handler

import (
	"context"
	"fmt"
	common "github.com/PonyWilliam/go-common"
	"github.com/micro/micro/v3/service/logger"
	"work/domain/model"
	work "work/domain/services"
	works "work/proto"
)

type Work struct{
	WorkService work.IWorkerServices
}
// Call is a single request handler called via client.Call or the generated client code
func(w *Work)CreateWorker(ctx context.Context,req *works.Request_Workers,res *works.Response_CreateWorker)error{
	workers := &model.Workers{}
	err := common.SwapTo(req,workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	id,err := w.WorkService.CreateWorker(workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	res.Id = id
	return nil
}
func(w *Work)UpdateWorker(ctx context.Context,req *works.Request_Workers,res *works.Response_CreateWorker)error{
	workers := &model.Workers{}
	err := common.SwapTo(req,workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	id,err := w.WorkService.CreateWorker(workers)
	if err != nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	res.Id = id
	return nil
}
func(w *Work)DeleteWorkerByID(ctx context.Context,req *works.Request_Workers_ID,res *works.Response_Workers)error{
	err := w.WorkService.DeleteWorkerByID(req.Id)
	if err!=nil{
		res.Message = err.Error()
		return err
	}
	res.Message = "success"
	return nil
}
func(w *Work)FindWorkerByID(ctx context.Context,req *works.Request_Workers_ID,res *works.Response_Worker_Show)error{
	worker,err := w.WorkService.FindWorkerByID(req.Id)
	fmt.Println(worker)
	workers := &works.Response_Workers_Info{}
	if err!=nil{
		return err
	}
	err = common.SwapTo(worker, workers)
	if err != nil{
		return err
	}
	res.Worker = workers
	return nil
}
func(w *Work)FindWorkerByName(ctx context.Context,req *works.Request_Workers_Name,res *works.Response_Workers_Show)error{
	workers,err := w.WorkService.FindWorkersByName(req.Name)
	if err != nil{
		return err
	}
	for _,v := range workers{
		worker := &works.Response_Workers_Info{}
		err = common.SwapTo(v,worker)
		res.Workers = append(res.Workers,worker)
	}
	if err != nil{
		return err
	}
	return nil
}
func(w *Work)FindAll(ctx context.Context,req *works.Request_Null,res *works.Response_Workers_Show)error{
	workers,err := w.WorkService.FindAll()
	if err != nil{
		return err
	}
	for _,v := range workers{
		worker := &works.Response_Workers_Info{}
		err = common.SwapTo(v,worker)
		res.Workers = append(res.Workers,worker)
	}
	if err != nil{
		return err
	}
	return nil
}
func(w *Work)CreateToken(ctx context.Context,req *works.LoginRequest,rsp *works.LoginResponse) error{
	ok,token,err := w.WorkService.Login(req.User,req.Password)
	if err!=nil || !ok{
		rsp.Code = ok
		rsp.Token = token
		logger.Error(err)
		logger.Error(ok)
		return err
	}
	rsp.Code = ok
	rsp.Token = token
	return nil
}

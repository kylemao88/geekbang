package handler

import (
	"fmt"
	gms "git.code.oa.com/gongyi/gomore/service"
	"git.code.oa.com/gongyi/yqz/pkg/common/logger" //这里引用了公司内网得公共包
	"git.code.oa.com/gongyi/yqz/pkg/common/util"   //这里引用了公司内网得公共包
	"google.golang.org/protobuf/proto"
	pb "study.geekbang.projs/api"
	fourmgr "study.geekbang.projs/internal/four"
	"time"
)

func QueryUserinfo(ctx *gms.Context) {
	defer util.GetUsedTime("QueryUserinfo")()
	var rsp *pb.QueryUserResponse
	defer func() {
		if rsp.Header.Code == 0 {
			rsp.Header.Msg = "success"
			logger.Info("%v - traceID: %v call suc, rsp: %v", util.GetCaller(), ctx.Request.TraceId, rsp)
		} else {
			logger.Info("%v - traceID: %v call failed, rsp: %v", util.GetCaller(), ctx.Request.TraceId, rsp)
		}
		rsp.Header.Msg += fmt.Sprintf(" traceID: %v", ctx.Request.TraceId)
		rsp.Header.OpTime = time.Now().Unix()
		err := ctx.Marshal(rsp)
		if err != nil {
			logger.Error("%v - ctx.Marshal rsp traceID: %v err: %s", util.GetCaller(), ctx.Request.TraceId, err.Error())
		}
	}()

	// decode request
	req := &pb.QueryUserRequest{}
	err := proto.Unmarshal(ctx.Request.Body, req)
	if err != nil {
		msg := fmt.Sprintf("%v - traceID: %v err: %s", util.GetCallee(), ctx.Request.TraceId, err.Error())
		logger.Error(msg)
		err = fmt.Errorf(msg)
		rsp = &pb.QueryUserResponse{
			Header: &pb.CommonHeader{
				Code: 99,
				Msg:  err.Error(),
			},
		}
		return
	}
	logger.Debug("%v - traceID: %v, svc: %v, req: %v", util.GetCallee(), ctx.Request.TraceId, ctx.Request.Method, req)

	user, err := fourmgr.QueryUserInfo(req.Uid)
	if err != nil {
		logger.Error("%v - traceID: %v err: %s", util.GetCallee(), ctx.Request.TraceId, err.Error())
		rsp = &pb.QueryUserResponse{
			Header: &pb.CommonHeader{
				Code: 99,
				Msg:  err.Error(),
			},
		}

		return
	}

	// response ...
	rsp = &pb.QueryUserResponse{
		Header: &pb.CommonHeader{
			Msg:  "success",
			Code: 0,
		},
		User: user,
	}
}

// Package main provides service entry functions
// 启动示例 1：
// ./activitymgr -busi_conf=../conf/ -namespace=Development -service=yqz.activitymgr -grpc_port=9001 -sidecar_port=19001 -http_port=8001
// 启动示例 2：
// ./activitymgr -busi_conf=../conf/ -namespace=Development -service=yqz.activitymgr -grpc_port=9001 -sidecar_port=19001 -http_port=8001 -enable_printscreen -log_level=5
//
// 对应的 sidecar 启动示例 1：
// ./sidecar -namespace=Development -service=yqz.activitymgr -http_port=18001 -grpc_port=19001 -svc_port=8001 -auto_register=false -skywalking=endpoint://127.0.0.1:11800 -sampling_rate=5000
// 对应的 sidecar 启动示例 2：
// ./sidecar -namespace=Development -service=yqz.activitymgr -http_port=18001 -grpc_port=19001 -svc_port=8001 -auto_register=false -skywalking=endpoint://127.0.0.1:11800 -sampling_rate=5000 -enable_printscreen -log_level=5

package main

import (
	"flag"
	"fmt"
	///  采用的是公司内微服务框架，内部通信采用GRPC服务
	gms "git.code.oa.com/gongyi/gomore/service"
	fourmgr "study.geekbang.projs/internal/four"
	"study.geekbang.projs/internal/handler"
)

var busiConf string

//命令行参数
func init() {
	// initialize global variable
	flag.StringVar(&busiConf, "busi_conf", "../conf/busi_conf.yaml", "Business configure dir path.")
}

func main() {
	flag.Parse()

	////  gms包是公司内微服务框架，暂未开源，外网可能暂时拉不到。。。额。。。
	svc := gms.NewService()
	svc.Handle("QueryUserinfo", handler.QueryUserinfo)
	gms.SimpleSvrMain(svc, checkParams, svrInit, svrFini)
}

// 参数校验
func checkParams() bool {
	if len(busiConf) == 0 {
		fmt.Println("Parameter[-busiConf] is not set")
		return false
	}
	return true
}

// 服务启动初始化
func svrInit() bool {
	return fourmgr.InitComponents(busiConf)
}

// 服务退出收尾
func svrFini() {
	//
}

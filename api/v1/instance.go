package v1

import (
	"go-skeleton/service"
	"go-skeleton/utils/log"

	"github.com/robfig/cron/v3"
)

var svc service.Service

func InitService() {
	svc = service.NewService()

	crontab := cron.New() // 默认从分开始进行时间调度
	// crontab := cron.New(cron.WithSeconds()) //精确到秒
	spec := "0 1 * * 0" //cron表达式，每周日凌晨1点更新
	crontab.AddFunc(spec, timedTask)
	// 启动定时器
	crontab.Start()
}

////////////////////////////////////////////////////////////////////////////////////
//
//      Internal function
//
////////////////////////////////////////////////////////////////////////////////////

func timedTask() {
	log.Infof("doing timed task")
}

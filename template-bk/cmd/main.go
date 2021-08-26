package main

import (
	"git.5th.im/lb-public/gear/cfg"
	"git.5th.im/lb-public/gear/log"
	microKit "git.5th.im/lb-public/gear/micro-kit"
	"git.5th.im/lb-public/gear/micro-kit/wrapper/logger"
	"git.5th.im/long-bridge-algo/golang/algo-dispatcher/pkg/metrics"
	"git.5th.im/long-bridge-algo/golang/algo-dispatcher/pkg/service"
	pb "git.5th.im/long-bridge-algo/golang/algo-dispatcher/proto/demo"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	svc := microKit.NewService(
		"lb.algo.demo",
		micro.WrapClient(logger.NewClientWrapper(
			logger.WithAllReply(true),
		)),
		micro.WrapHandler(logger.NewServerWrapper(
			logger.WithAllReply(true),
		)),
		micro.Flags(
			cli.StringFlag{
				Name:  "lvl",
				Usage: "log level",
				Value: "info",
			},
			cli.StringFlag{
				Name:  "config",
				Usage: "service config path",
			},
		),
	)

	var hd *service.Handler

	svc.Init(
		micro.Action(func(c *cli.Context) {
			log.SetLevel(log.ParseLevel(c.String("lvl")))
			// 加载配置
			cfg.LoadDirs("./config/common", c.String("config"))

			hd = service.NewHandler(svc.Client())

			// 注册健康检查的资源
			svc.RegisterHealthWatchList(hd)
			// 注册 Handlers
			_ = pb.RegisterAlgoDispatcherHandler(svc.Server(), hd)
		}),

		micro.AfterStop(func() error {
			return nil
		}),
	)

	// 注册 prometheus metrics 监控
	metrics.Register()
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}

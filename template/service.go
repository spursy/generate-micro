package template

var (
	ServiceSRV = `package service

import (
	"context"
	"encoding/json"
	"errors"
	"git.5th.im/lb-public/gear/cfg"
	microKit "git.5th.im/lb-public/gear/micro-kit"
	bs "{{.Dir}}/pkg/business"
	"{{.Dir}}/pkg/metrics"
	pb "{{.Dir}}/proto/demo"
	"github.com/demdxx/gocast"
	"github.com/micro/go-micro/client"
	"log"
)

type Handler struct {
	business *bs.Business
}

// 健康检查
func (h *Handler) HealthWatchList() []microKit.Healthy {
	return []microKit.Healthy{h.business.Redis}
}

func NewHandler(cli client.Client) *Handler {
	r := cfg.NewRedis("redis")
	handler := Handler{
		business: bs.New(r),
	}

	return &handler
}

// 断开数据库链接
func (h *Handler) Close() error {
	_ = h.business.Redis.Close()
	return nil
}

// Demo Func
func (h *Handler) GetDemo(ctx context.Context, req *pb.DemoRequest, resp *pb.DemoListReply) error {
	emptyList := make([]*pb.DemoObject, 0)
	defer func() {
		var customError error
		if err := recover(); err != nil {
			customError = errors.New(gocast.ToString(err))
			log.Printf("The error is %v\n", err)
			resp.DemoList = emptyList
		}
		metrics.CountTask("GetDemo", "AlgoDispatcher", customError)
	}()
	measure := metrics.MeasureMethod("AlgoDispatcher", "GetDemo")
	defer measure.Done()

	// 打印请求参数
	byteJson, _ := json.Marshal(req)
	log.Printf("req params is: %s\n", byteJson)

	params := make(map[string]interface{})
	params["k"] = []string{req.K}

	demoList, err := h.business.GetDemo(ctx, params)

	// 如果有异常
	if err != nil {
		log.Printf("The error is %v\n", err)
		resp.DemoList = emptyList
		metrics.CountTask("GetDemo", "AlgoDispatcher", err)
	} else {
		log.Printf("The length is: %v\n", len(demoList))
		resp.DemoList = demoList
	}
	return nil
}

`
)




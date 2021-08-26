package business

import (
	"context"
	"fmt"
	"git.5th.im/lb-public/gear/cache"
	pb "git.5th.im/long-bridge-algo/golang/algo-dispatcher/proto/demo"
)

/**
Search 结构体
*/
type Business struct {
	Redis *cache.RedisClient // redis client db
}

func (bs *Business) GetDemo(ctx context.Context, params map[string]interface{}) ([]*pb.DemoObject, error) {
	fmt.Printf("params is %+v\n", params)
	demoList := make([]*pb.DemoObject, 0)

	demoList = append(demoList, &pb.DemoObject{
		Id:    1,
		Title: "Test",
	})
	return demoList, nil
}

/**
构建 Search 结构体
*/
func New(redis *cache.RedisClient) *Business {
	business := &Business{
		Redis: redis,
	}
	return business
}

/*
断开链接
*/
func (bs *Business) Close() error {
	_ = bs.Redis.Close()
	return nil
}

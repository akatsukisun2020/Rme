package dao

import (
	"context"
	"encoding/json"
	"rme/internal/model"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type loginInfoRedisClient struct {
	redis *gredis.Redis
}

func NewLoginInfoRedisClient() *loginInfoRedisClient {
	return &loginInfoRedisClient{
		redis: g.Redis(),
	}
}

func makeLoginInfoRedisKey(userID string) string {
	return "rme_logininfo_" + userID
}

// TODO: 增加过期时间的概念，另外，测试一下redis的性能？最大连接数等逻辑？
func (cli *loginInfoRedisClient) Set(ctx context.Context, info *model.LoginInfo) error {
	key := makeLoginInfoRedisKey(info.UserId)
	value, _ := json.Marshal(info)

	_, err := cli.redis.Set(ctx, key, string(value))
	if err != nil {
		g.Log().Errorf(ctx, "loginInfoRedisClient.Set error, key:%s, err:%v", key, err)
		return err
	}

	g.Log().Debugf(ctx, "loginInfoRedisClient.Set success, key:%s", key)
	return nil
}

func (cli *loginInfoRedisClient) Get(ctx context.Context, userID string) (*model.LoginInfo, error) {
	key := makeLoginInfoRedisKey(userID)

	value, err := cli.redis.Get(ctx, key)
	if err != nil {
		g.Log().Errorf(ctx, "loginInfoRedisClient.Get error, key:%s, err:%v", key, err)
		return nil, err
	}

	info := new(model.LoginInfo)
	if err = json.Unmarshal(value.Bytes(), info); err != nil {
		g.Log().Errorf(ctx, "loginInfoRedisClient.Unmarshal error, key:%s, err:%v", key, err)
		return nil, err
	}

	g.Log().Debugf(ctx, "loginInfoRedisClient.Get success, key:%s, info:%v", key, info)
	return info, nil
}

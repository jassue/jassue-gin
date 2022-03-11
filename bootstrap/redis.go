package bootstrap

import (
    "context"
    "github.com/go-redis/redis/v8"
    "github.com/jassue/jassue-gin/global"
    "go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
        Password: global.App.Config.Redis.Password, // no password set
        DB:       global.App.Config.Redis.DB,       // use default DB
    })
    _, err := client.Ping(context.Background()).Result()
    if err != nil {
        global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
        return nil
    }
    return client
}

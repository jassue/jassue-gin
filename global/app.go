package global

import (
    "github.com/go-redis/redis/v8"
    "github.com/jassue/go-storage/storage"
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "gorm.io/gorm"
    "jassue-gin/config"
)

type Application struct {
    ConfigViper *viper.Viper
    Config config.Configuration
    Log *zap.Logger
    DB *gorm.DB
    Redis *redis.Client
}

var App = new(Application)

func (app *Application) Disk(disk... string) storage.Storage {
    diskName := app.Config.Storage.Default
    if len(disk) > 0 {
        diskName = storage.DiskName(disk[0])
    }
    s, err := storage.Disk(diskName)
    if err != nil {
        panic(err)
    }
    return s
}

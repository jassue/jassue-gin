package bootstrap

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/jassue/jassue-gin/global"
    "github.com/spf13/viper"
    "os"
)

func InitializeConfig() *viper.Viper {
    // 设置配置文件路径
    config := "config.yaml"
    if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
        config = configEnv
    }

    // 初始化 viper
    v := viper.New()
    v.SetConfigFile(config)
    v.SetConfigType("yaml")
    if err := v.ReadInConfig(); err != nil {
        panic(fmt.Errorf("read config failed: %s \n", err))
    }

    // 监听配置文件
    v.WatchConfig()
    v.OnConfigChange(func(in fsnotify.Event) {
        fmt.Println("config file changed:", in.Name)
        // 重载配置
        if err := v.Unmarshal(&global.App.Config); err != nil {
            fmt.Println(err)
        }
    })
    // 将配置赋值给全局变量
    if err := v.Unmarshal(&global.App.Config); err != nil {
       fmt.Println(err)
    }

    return v
}

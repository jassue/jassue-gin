package bootstrap

import (
    "fmt"
    "github.com/jassue/jassue-gin/global"
    "github.com/robfig/cron/v3"
    "time"
)

func InitializeCron() {
    global.App.Cron = cron.New(cron.WithSeconds())

    go func() {
       global.App.Cron.AddFunc("0 0 2 * * *", func() {
           fmt.Println(time.Now())
       })
       global.App.Cron.Start()
       defer global.App.Cron.Stop()
       select {}
    }()
}

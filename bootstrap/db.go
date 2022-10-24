package bootstrap

import (
    "github.com/jassue/jassue-gin/app/models"
    "github.com/jassue/jassue-gin/global"
    "go.uber.org/zap"
    "gopkg.in/natefinch/lumberjack.v2"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "io"
    "log"
    "os"
    "strconv"
    "time"
)

func InitializeDB() *gorm.DB {
    // 根据驱动配置进行初始化
    switch global.App.Config.Database.Driver {
    case "mysql":
       return initMySqlGorm()
    default:
       return initMySqlGorm()
    }
}

func initMySqlGorm() *gorm.DB {
    dbConfig := global.App.Config.Database

    if dbConfig.Database == "" {
        return nil
    }
    dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
        dbConfig.Database + "?charset=" + dbConfig.Charset +"&parseTime=True&loc=Local"
    mysqlConfig := mysql.Config{
        DSN:                       dsn,   // DSN data source name
        DefaultStringSize:         191,   // string 类型字段的默认长度
        DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
        DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
        DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
        SkipInitializeWithVersion: false, // 根据版本自动配置
    }
    if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
        Logger: getGormLogger(), // 使用自定义 Logger
    }); err != nil {
        global.App.Log.Error("mysql connect failed, err:", zap.Any("err", err))
        return nil
    } else {
        sqlDB, _ := db.DB()
        sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
        sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
        initMySqlTables(db)
        return db
    }
}

func getGormLogger() logger.Interface {
    var logMode logger.LogLevel

    switch global.App.Config.Database.LogMode {
    case "silent":
        logMode = logger.Silent
    case "error":
        logMode = logger.Error
    case "warn":
        logMode = logger.Warn
    case "info":
        logMode = logger.Info
    default:
        logMode = logger.Info
    }

    return logger.New(getGormLogWriter(), logger.Config{
        SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
        LogLevel:                  logMode, // 日志级别
        IgnoreRecordNotFoundError: false, // 忽略ErrRecordNotFound（记录未找到）错误
        Colorful:                  !global.App.Config.Database.EnableFileLogWriter, // 禁用彩色打印
    })
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
    var writer io.Writer

    // 是否启用日志文件
    if global.App.Config.Database.EnableFileLogWriter {
        // 自定义 Writer
        writer = &lumberjack.Logger{
            Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFilename,
            MaxSize:    global.App.Config.Log.MaxSize,
            MaxBackups: global.App.Config.Log.MaxBackups,
            MaxAge:     global.App.Config.Log.MaxAge,
            Compress:   global.App.Config.Log.Compress,
        }
    } else {
        // 默认 Writer
        writer = os.Stdout
    }
    return log.New(writer, "\r\n", log.LstdFlags)
}

// 数据库表初始化
func initMySqlTables(db *gorm.DB) {
    err := db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(
        models.User{},
        models.Media{},
    )
    if err != nil {
        global.App.Log.Error("migrate table failed", zap.Any("err", err))
        os.Exit(0)
    }
}

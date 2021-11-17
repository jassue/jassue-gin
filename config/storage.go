package config

import (
    "github.com/jassue/go-storage/kodo"
    "github.com/jassue/go-storage/local"
    "github.com/jassue/go-storage/oss"
    "github.com/jassue/go-storage/storage"
)

type Storage struct {
    Default storage.DiskName `mapstructure:"default" json:"default" yaml:"default"` // local本地 oss阿里云 kodo七牛云
    Disks Disks `mapstructure:"disks" json:"disks" yaml:"disks"`
}

type Disks struct {
    Local local.Config `mapstructure:"local" json:"local" yaml:"local"`
    AliOss oss.Config `mapstructure:"ali_oss" json:"ali_oss" yaml:"ali_oss"`
    QiNiu kodo.Config `mapstructure:"qi_niu" json:"qi_niu" yaml:"qi_niu"`
}

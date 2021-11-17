package models

type Media struct {
    ID
    DiskType string `json:"disk_type" gorm:"size:20;index;not null;comment:存储类型"`
    SrcType int8 `json:"src_type" gorm:"not null;comment:链接类型 1相对路径 2外链"`
    Src string `json:"src" gorm:"not null;comment:资源链接"`
    Timestamps
}

package models

import "strconv"

type User struct {
    ID
    Name string `json:"name" gorm:"size:30;not null;comment:用户名称"`
    Mobile string `json:"mobile" gorm:"size:24;not null;index;comment:用户手机号"`
    Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
    Timestamps
    SoftDeletes
}

func (user User) GetUid() string {
    return strconv.Itoa(int(user.ID.ID))
}

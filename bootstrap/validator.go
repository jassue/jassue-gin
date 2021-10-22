package bootstrap

import (
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
    "jassue-gin/utils"
    "reflect"
    "strings"
)

func InitializeValidator() {
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        // 注册自定义验证器
        _ = v.RegisterValidation("mobile", utils.ValidateMobile)

        // 注册自定义 json tag 函数
        v.RegisterTagNameFunc(func(fld reflect.StructField) string {
            name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
            if name == "-" {
                return ""
            }
            return name
        })
    }
}

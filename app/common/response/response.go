package response

import (
    "github.com/gin-gonic/gin"
    "jassue-gin/global"
    "net/http"
)

type Response struct {
    ErrorCode int `json:"error_code"`
    Data interface{} `json:"data"`
    Message string `json:"message"`
}

func Success(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, Response{
        0,
        data,
        "ok",
    })
}

func Fail(c *gin.Context, errorCode int, msg string) {
    c.JSON(http.StatusOK, Response{
        errorCode,
        nil,
        msg,
    })
}

func FailByError(c *gin.Context, error global.CustomError) {
    Fail(c, error.ErrorCode, error.ErrorMsg)
}

func ValidateFail(c *gin.Context, msg string)  {
    Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

func BusinessFail(c *gin.Context, msg string) {
    Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}

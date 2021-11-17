package common

import (
    "github.com/gin-gonic/gin"
    "jassue-gin/app/common/request"
    "jassue-gin/app/common/response"
    "jassue-gin/app/services"
)

func ImageUpload(c *gin.Context) {
    var form request.ImageUpload
    if err := c.ShouldBind(&form); err != nil {
        response.ValidateFail(c, request.GetErrorMsg(form, err))
        return
    }

    outPut, err := services.MediaService.SaveImage(form)
    if err != nil {
        response.BusinessFail(c, err.Error())
        return
    }
    response.Success(c, outPut)
}

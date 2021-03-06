package handler

import (
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
	登陆
*/
func Login(ctx *gin.Context) {
	//获取参数
	req := &request.LoginReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	resp, err := service.LoginSrv(ctx, req)
	if err != nil {
		log.Errorf("Login resp is error (%v)", err)
		Fail(ctx, enums.ErrorSystemException)
		return
	}
	Success(ctx, resp)
}

/*
	退出
*/
func Logout(ctx *gin.Context) {
	//获取参数
	req := &request.LogoutReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	err := service.LogoutSrv(ctx, req)
	if err != nil {
		log.Errorf("Logout resp is error (%v)", err)
		Fail(ctx, enums.ErrorSystemException)
		return
	}
	Success(ctx, nil)
}

/*
	图片上传
*/
func ImageUpload(ctx *gin.Context) {
	//获取参数
	req := &request.ImageUploadReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	resp, err := service.ImageUploadSrv(ctx, req)
	if err != nil {
		log.Errorf("ImageUpload resp is error (%v)", err)
		Fail(ctx, enums.ErrorSystemException)
		return
	}
	Success(ctx, resp)
}

/*
	添加枚举
*/
func AddEnum(ctx *gin.Context) {
	//获取参数
	req := &request.AddEnumReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	err := service.AddEnumSrv(ctx, req)
	if err != nil {
		log.Errorf("AddEnum is error (%v)", err)
		Fail(ctx, enums.ErrorSystemException)
		return
	}
	Success(ctx, nil)
}

/*
	删除枚举
*/
func DelEnum(ctx *gin.Context) {
	//获取参数
	req := &request.DelEnumReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	err := service.DelEnumSrv(ctx, req)
	if err != nil {
		log.Errorf("AddEnum is error (%v)", err)
		Fail(ctx, enums.ErrorSystemException)
		return
	}
	Success(ctx, nil)
}

/*
	获取枚举
*/
func GetEnumList(ctx *gin.Context) {
	//获取参数
	req := &request.GetEnumReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	resp, err := service.GetEnumSrv(ctx, req)
	if err != nil {
		log.Errorf("AddEnum is error (%v)", err)
		Fail(ctx, enums.ErrorSystemException)
		return
	}
	Success(ctx, resp)
}

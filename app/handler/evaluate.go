package handler

import (
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

/*
	产品list
*/
func GetProductList(ctx *gin.Context) {
	//获取参数
	req := &request.GetProductListReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	resp, err := service.GetProductListSrv(ctx, req)
	if err != nil {
		log.Errorf("GetProductList resp is error (%v)", err)
		Fail(ctx, enums.ErrorSystemException)
		return
	}
	Success(ctx, resp)
}

/*
	产品范围list
*/
func GetProductRangeList(ctx *gin.Context) {
	//获取参数
	req := &request.GetProductRangeListReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	resp, err := service.GetProductRangeListSrv(ctx, req)
	if err != nil {
		log.Errorf("GetProductList resp is error (%v)", err)
		Fail(ctx, enums.ErrorSystemException)
		return
	}
	Success(ctx, resp)
}

/*
	产品详情
*/
func GetProductInfo(ctx *gin.Context) {
	//获取参数
	req := &request.GetProductInfoReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	resp, err := service.GetProductInfoSrv(ctx, req)
	if err != nil {
		log.Errorf("GetProductInfo resp is error (%v)", err)
		Fail(ctx, enums.ErrorSystemException)
		return
	}
	Success(ctx, resp)
}

/*
	编辑产品
*/
func EditProduct(ctx *gin.Context) {
	//获取参数
	req := &request.EditProductReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	err := service.EditProductSrv(ctx, req)
	if err != nil {
		Fail(ctx, enums.ErrorFileUpdateFail)
		return
	}
	Success(ctx, nil)
}

/*
	删除产品
*/
func DelProduct(ctx *gin.Context) {
	//获取参数
	req := &request.DelProductReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	err := service.DelProductSrv(ctx, req)
	if err != nil {
		Fail(ctx, enums.ErrorFileUpdateFail)
		return
	}
	Success(ctx, nil)
}

/*
	图片下载
*/
func ImageDownload(ctx *gin.Context) {
	//获取参数
	req := &request.ImageDownloadReq{}
	if err := req.Validate(ctx); err != nil {
		log.Errorf("params validate error (%v)", err)
		Fail(ctx, enums.ErrorInputValidate)
		return
	}
	err := service.ImageDownloadSrv(ctx, req)
	if err != nil {
		Fail(ctx, enums.ErrorSystemException)
		return
	}
}

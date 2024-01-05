package grreport

import (
	"shservice/common"
	"shservice/crv"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

type GRReportController struct {
	CRVClient *crv.CRVClient
}

func (controller *GRReportController)downloadPRPSReport(c *gin.Context){
	log.Println("GRReportController start downloadPRPSReport")
	
	var header crv.CommonHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		log.Println(err)
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		log.Println("end downloadPRPSReport with error")
		return
	}	
	
	var rep crv.CommonReq
	if err := c.BindJSON(&rep); err != nil {
		log.Println(err)
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		return
  }

	if rep.SelectedRowKeys==nil || len(*rep.SelectedRowKeys)==0 {
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		log.Println("error：request list is empty")
		return
	}

	//获取数据
	res:=GetGRPSReportData(rep.SelectedRowKeys,controller.CRVClient,header.Token)
	log.Println(res)

	if res!=nil && len(res)>0 {
		var fileName string
		if len(res)==1 {
			fileName=GetReportFileName(res[0].(map[string]interface{}))+".xlsx"
		} else {
			fileName=GetBatchID()+".zip"
		}

		fileName=url.QueryEscape(fileName)
		
		c.Header("Content-Type", "application/octet-stream")
    c.Header("Content-Disposition", "attachment; filename="+fileName)
    c.Header("Content-Transfer-Encoding", "binary")
	
		//生成报表
		CreateExcelReports("./templetes","grps",res,c.Writer)
	}
	
	log.Println("GRReportController end downloadPRPSReport")
}

func (controller *GRReportController)downloadPRMSReport(c *gin.Context){
	log.Println("GRReportController start downloadPRMSReport")
	
	var header crv.CommonHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		log.Println(err)
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		log.Println("end downloadPRMSReport with error")
		return
	}	
	
	var rep crv.CommonReq
	if err := c.BindJSON(&rep); err != nil {
		log.Println(err)
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		return
  	}	

	if rep.SelectedRowKeys==nil || len(*rep.SelectedRowKeys)==0 {
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		log.Println("error：request list is empty")
		return
	}

	//获取数据
	res:=GetGRMSReportData(rep.SelectedRowKeys,controller.CRVClient,header.Token)
	log.Println(res)

	if res!=nil && len(res)>0 {
		var fileName string
		if len(res)==1 {
			fileName=GetReportFileName(res[0].(map[string]interface{}))+".pdf"
		} else {
			fileName=GetBatchID()+".zip"
		}

		fileName=url.QueryEscape(fileName)
		
		c.Header("Content-Type", "application/octet-stream")
    c.Header("Content-Disposition", "attachment; filename="+fileName)
    c.Header("Content-Transfer-Encoding", "binary")
	
		//生成报表
		tmpName:=GetReportTemplete("grms",res[0].(map[string]interface{}))
		CreateReports(tmpName,res,c.Writer)
	}
	
	log.Println("GRReportController end downloadPRMSReport")
}

//Bind bind the controller function to url
func (controller *GRReportController) Bind(router *gin.Engine) {
	log.Println("Bind GRReportController")
	router.POST("/downloadPRPSReport", controller.downloadPRPSReport)
	router.POST("/downloadPRMSReport", controller.downloadPRMSReport)
}
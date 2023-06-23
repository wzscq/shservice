package main

import (
	"testing"
	"shservice/grreport"
	"log"
)

func TestGRMSTemplete(t *testing.T) {
	//加载模板
	tmp,err:=grreport.LoadTempleteFromJsonFile("./templetes/grms.json")
	if err!=nil {
		log.Println(err)
		return
	}
	//创建报告临时存放目录
	outFile:="./reportfile/output.pdf"

	pdf:=grreport.CreatePdf(tmp,nil)
	pdf.WritePdf(outFile)
}

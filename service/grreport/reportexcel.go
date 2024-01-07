package grreport

import (
	"shservice/common"
	"archive/zip"
	"io"
	"log"
	"strconv"
	"fmt"
	"github.com/wzscq/exceltemplate"
)

func CreateExcelReports(templatePath,tmpType string,list []interface{},w io.Writer)(int){
	//如果下载的报告数量大于1，则打包成zip
	var zipWriter *zip.Writer
	if len(list)>1 {
		zipWriter= zip.NewWriter(w)
		defer zipWriter.Close()
	}

	for _,data:=range list {
		mapData:=data.(map[string]interface{})
		getTemplateFileName:=getTemplateFileName(tmpType,mapData)
		//检查模板文件是否存在
		excel,err:=exceltemplate.GetExcelFromTemplate(templatePath+"/"+getTemplateFileName,mapData)
		if err!=nil {
			return common.ResultReadTempleteFileError
		}

		if zipWriter!=nil {
			outFileName:=GetReportFileName(mapData)+".xlsx"
			fileHeader:=&zip.FileHeader{
				Name:outFileName,
			}
			fileWriter,err:=zipWriter.CreateHeader(fileHeader)
			if err!=nil {
				log.Println(err)
			} 
			excel.Write(fileWriter)
		} else {
			excel.Write(w)
		}
		excel.Close()		
	}

	return common.ResultSuccess
}

func getTemplateFileName(tmpType string,data map[string]interface{})(string){
	//模板命名规则：x年级_x学期.xlsx
	semesterModel:=data["semester"].(map[string]interface{})
	semester:=semesterModel["value"].(string)
	year:=data["year"].(string)
	classModel:=data["class"].(map[string]interface{})
	classList:=classModel["list"].([]interface{})
	class:=classList[0].(map[string]interface{})
	enrollment_year:=class["enrollment_year"].(string)
	//当前年级是year-enrollment_year+1
	yearInt,_:=strconv.Atoi(year)
	enrollment_yearInt,_:=strconv.Atoi(enrollment_year)
	grade:=yearInt-enrollment_yearInt+1
	return fmt.Sprintf("%s_%d年级_%s学期.xlsx",tmpType,grade,semester)
}



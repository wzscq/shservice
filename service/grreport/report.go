package grreport

import (
	"shservice/common"
	"archive/zip"
	"os"
	"io"
	"log"
)

func CreateReports(tmpName string,list []interface{},w io.Writer)(int){
	//加载模板
	tmp,err:=LoadTempleteFromJsonFile("./templetes/"+tmpName+".json")
	if err!=nil {
		return common.ResultReadTempleteFileError
	}
	//创建报告临时存放目录
	outPath:="./reportfile/"+GetBatchID();
	CreateDir(outPath)

	//如果下载的报告数量大于1，则打包成zip
	var zipWriter *zip.Writer
	if len(list)>1 {
		zipWriter= zip.NewWriter(w)
		defer zipWriter.Close()
	}

	for _,data:=range list {
		mapData:=data.(map[string]interface{})
		pdf:=CreatePdf(tmp,mapData)
		if pdf!=nil {
			if zipWriter!=nil {
				outFileName:=GetReportFileName(mapData)+".pdf"
				fileHeader:=&zip.FileHeader{
					Name:outFileName,
				}
				fileWriter,err:=zipWriter.CreateHeader(fileHeader)
				if err!=nil {
					log.Println(err)
				} 
				pdf.Write(fileWriter)
			} else {
			
				pdf.Write(w)
			}
		}
	}

	return common.ResultSuccess
}

func GetReportFileName(data map[string]interface{})(string){
	//学年_学期_学号_姓名
	year:=data["year"].(string)
	semesterModel:=data["semester"].(map[string]interface{})
	semester:=semesterModel["value"].(string)
	studentModel:=data["student"].(map[string]interface{})
	studentList:=studentModel["list"].([]interface{})
	student:=studentList[0].(map[string]interface{})
	sn:=student["sn"].(string)
	name:=student["name"].(string)
	return year+"_"+semester+"_"+sn+"_"+name
}

func CreateDir(dir string)(error){
	//创建目录
	err:=os.MkdirAll(dir,os.ModePerm)
	return err
}
